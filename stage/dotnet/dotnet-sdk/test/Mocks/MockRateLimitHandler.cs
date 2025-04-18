// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

namespace Tests.Mocks;

public class MockRateLimitHandler : HttpMessageHandler
{
    private HttpResponseMessage? Response1
    {
        get; set;
    }

    private HttpResponseMessage? Response2
    {
        get; set;
    }

    private bool _response1Sent = false;

    protected override async Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
    {
        if (!_response1Sent)
        {
            _response1Sent = true;
            if (Response1 != null)
            {
                Response1.RequestMessage = request;
                return await Task.FromResult(Response1);
            }
            else
            {
                return await Task.FromResult(new HttpResponseMessage());
            }
        }
        else
        {
            _response1Sent = false;
            if (Response2 != null)
            {
                Response2.RequestMessage = request;
                return await Task.FromResult(Response2);
            }
            else
            {
                return await Task.FromResult(new HttpResponseMessage());
            }
        }
    }

    public void SetHttpResponse(HttpResponseMessage response1, HttpResponseMessage? response2 = null)
    {
        _response1Sent = false;
        Response1 = response1;
        Response2 = response2;
    }
}
