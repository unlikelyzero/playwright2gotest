package test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestCrowdstrike(t *testing.T) {
	client := &http.Client{}
	var err error
	var req *http.Request
	var resp *http.Response

	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/en-us/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/consent/bee15b7c-b632-450e-9003-9c8b60b3b978/OtAutoBlock.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/scripttemplates/otSDKStub.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/launch-d9bfd4283ab8.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-analytics.lc-f643ec19011038440608ac5e1d5dd675-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-grid.lc-d75b4ada966a12a0acc4f4483174f716-lc.min.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-base.lc-61e1f474009a3adcff6cc95dde2309a3-lc.min.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-language.lc-4418ee9ad99afc3c80f14f47d372a699-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts.lc-475459b8ddb225de7aad82eaf0d465c8-lc.min.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/clientlibs/granite/jquery.lc-f9e8e8c279baf6a1a278042afe4f395a-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/clientlibs/granite/utils.lc-899004cc02c33efc1f6694b1aee587fd-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-dependencies.lc-bb1dbdd53e32240429436d11ecb5f036-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-dependencies.lc-2929dd9d69a653da2cf8fe429017e6b6-lc.min.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-common.lc-d2017b1bb6b94dc55ac836b3e60fc92d-lc.min.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-header.lc-0b9b1c99ad703dffc489ae2fdecad2b2-lc.min.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-addsearch.lc-fa338b27fbcf1b58755f6bc3391495d1-lc.min.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-dotcom.lc-33079283bcb9b23ffc4ce228005513c1-lc.min.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/graphics-and-illustrations/marketplace/addsearch-icon.svg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/login-icon", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/shopping-cart-empty", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/.rum/@adobe/helix-rum-js@%5E2/src/index.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/empty-cart-image", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/experience-breach-Icon", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/phone-icon", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://play.vidyard.com/72xzgRJUfHEGrXeQaYMY5M.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/adversaries/open-report.png", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://play.vidyard.com/PBjCLnZ26eK9VG3ifwnnVp.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://play.vidyard.com/SJBQ6ji2tUz4syr4dz2y3s.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://play.vidyard.com/nTrNLYDyFzyH9dBgveMv91.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/adversaries/hi-res-adv-bg.png", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/adversaries/adversary-group.png", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/privacyoptions-icon-footer", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-common.lc-9bb6aad5b81df0110105b0a332fb1ecd-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-header.lc-13360fa8c738773238c6e841d8970e35-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-addsearch.lc-b8f6141a81735ad2112829f5c5025005-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-dotcom.lc-db18d2853624d228ab184335c97a1302-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/core/wcm/components/commons/site/clientlibs/container.lc-0a6aff292f5cc42142779cde92054524-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-base.lc-98b44ec74775c5bc76b0744df1c9b66c-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-lottie.lc-59169d2e9977544fd3ae80d885311bce-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-animation.lc-012003c5a08ba65bfd2d6d6b777b7f4e-lc.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/consent/bee15b7c-b632-450e-9003-9c8b60b3b978/bee15b7c-b632-450e-9003-9c8b60b3b978.json", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://epsilon.6sense.com/v3/company/details", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://use.typekit.net/zya3koo.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://p.typekit.net/p.css?s=1&amp;amp;k=zya3koo&amp;amp;ht=tk&amp;amp;f=39496.39498.39500&amp;amp;a=30979937&amp;amp;app=typekit&amp;amp;e=css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.vidyard.com/thumbnails/47110555/-S4aN8AbSJIrGOuNu7CWYQ.png", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.vidyard.com/thumbnails/40870997/dRaHCM-NsYYKXH0xN7hZDnL47TXyyc1z.gif", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.vidyard.com/thumbnails/37747693/PZg_du-4KYqqTO52YjfSplzbDO2VXLBE.gif", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.vidyard.com/thumbnails/40480664/eVCtvV3j2pznjltqlZVCCQ6gHkn4tKge.gif", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/backgrounds/falcon-platform-texture-bg-v2.png", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/haas_grot_disp/HaasGrotDisp-65Medium.woff2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/neue_haas_grot_disp_pro/NeueHaasDisplayLight.ttf", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/sans/Crowdstrike_Sharp_Sans/CrowdstrikeSharpSans-Mdm.woff2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/neue_haas_grot_disp_pro/NeueHaasDisplayMedium.ttf", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/sans/Crowdstrike_Sharp_Sans/CrowdstrikeSharpSans-Bold.woff2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/haas_grot_disp/HaasGrotDisp-55Roman.woff2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/haas_grot_disp/HaasGrotDisp-45Light.woff2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/black-primary-crowdstrike-logo-1-addedPadding-3?ts=1746646455211&amp;amp;dpr=off", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Platform-4-16-25-darkmode-HP?ts=1746621981798&amp;amp;dpr=off", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/image/crowdstrikeinc/25-SRV-007_Forrester%20Wave%20MDR%203153x3860?ts=1746621981959&amp;amp;dpr=off", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/image/crowdstrikeinc/cloud-lock-generic-img?ts=1746621982750&amp;amp;dpr=off", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/image/crowdstrikeinc/identity-security-services-new?ts=1746621982825&amp;amp;dpr=off", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/image/crowdstrikeinc/soc-survival-guide-2.bak?ts=1746621982899&amp;amp;dpr=off", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/scripttemplates/202401.2.0/otBannerSdk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/snowflake/analyticsdata?timestamp=1746729078679", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;event-name&#34;:&#34;vieworder&#34;,&#34;tag&#34;:&#34;link&#34;,&#34;tag-location&#34;:&#34;&#34;,&#34;landingpage-url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;window-inner-width&#34;:1280,&#34;window-inner-height&#34;:720,&#34;body-client-width&#34;:1280,&#34;body-client-height&#34;:720,&#34;screen-width&#34;:1280,&#34;screen-height&#34;:720,&#34;screen-width-height&#34;:&#34;1280x720&#34;,&#34;device-pixel-ratio&#34;:1,&#34;browser-name&#34;:&#34;Chrome&#34;,&#34;browser-version&#34;:&#34;136.0.7103.25&#34;,&#34;os-name&#34;:&#34;Windows&#34;,&#34;os-version&#34;:&#34;NT 10.0&#34;,&#34;engine-name&#34;:&#34;Blink&#34;,&#34;engine-version&#34;:&#34;&#34;,&#34;device-type&#34;:&#34;desktop&#34;,&#34;session_id&#34;:&#34;8ffaa7ae-9b97-463d-b437-ce63a0483042&#34;,&#34;productName&#34;:&#34;&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/image/crowdstrikeinc/try-free-footer", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/disableDomains", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/experience-fragments/crowdstrike-www/locale-sites/us/en-us/site/header/header-elements/platform/platform-redesign.content.html", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/experience-fragments/crowdstrike-www/locale-sites/us/en-us/site/header/header-elements/services/services-redesign.content.html", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/experience-fragments/crowdstrike-www/locale-sites/us/en-us/site/header/header-elements/Solutions/master.content.html", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/experience-fragments/crowdstrike-www/locale-sites/us/en-us/site/header/header-elements/why-crowdstrike/why-crowdstrike-redesign.content.html", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/content/experience-fragments/crowdstrike-www/locale-sites/us/en-us/site/header/header-elements/Resources/master.content.html", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/image/crowdstrikeinc/charlotte-hp-hero-1?&amp;amp;hei=769&amp;amp;wid=1600&amp;amp;qlt=90", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/sans/Crowdstrike_Sharp_Sans/CrowdstrikeSharpSans-Mdm.ttf", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/sans/Crowdstrike_Sharp_Sans/CrowdstrikeSharpSans-Bold.ttf", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/image/crowdstrikeinc/gartner-mq-leader-report?ts=1746621982064&amp;amp;dpr=off", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/image/crowdstrikeinc/frost-radar-2024-2?ts=1746621982168&amp;amp;dpr=off", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/consent/bee15b7c-b632-450e-9003-9c8b60b3b978/e707ec79-97bc-46fd-8adb-adb5851b7f24/en.json", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Endpoint-Security-2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Exposure-Management-2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Generative-AI", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Identity-Protection-1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/SaaS-Security", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/IT-Automation", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Threat-Intelligence-Hunting", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Cloud-Security-2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Workflow-Automation", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Next-Gen-SIEM", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/Data-Protection-3", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.crowdstrike.com/is/content/crowdstrikeinc/xiot-security", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/scripttemplates/202401.2.0/assets/otFlat.json", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/scripttemplates/202401.2.0/assets/v2/otPcTab.json", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/scripttemplates/202401.2.0/assets/otCommonStyles.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/launch-d9bfd4283ab8.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/logos/static/ot_close.svg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/logos/static/ot_guard_logo.svg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/logos/c109dae9-46f3-4e91-a59e-7844ef645107/a4486e7b-87d4-4354-a844-b53722d937ac/e2865569-21ec-4213-8a1d-4f0ed5a1672e/CS_Logo_2022_In-Line_All-Red_RGB_(1).png", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.cookielaw.org/logos/static/powered_by_logo.svg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://connect.facebook.net/en_US/fbevents.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://adobedc.demdex.net/ee/v1/interact?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=35dcc33a-51ce-4b66-bd7c-75459093b347", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;events&#34;:[{&#34;xdm&#34;:{&#34;eventType&#34;:&#34;decisioning.propositionFetch&#34;,&#34;web&#34;:{&#34;webPageDetails&#34;:{&#34;URL&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;},&#34;webReferrer&#34;:{&#34;URL&#34;:&#34;&#34;}},&#34;device&#34;:{&#34;screenHeight&#34;:720,&#34;screenWidth&#34;:1280,&#34;screenOrientation&#34;:&#34;landscape&#34;},&#34;environment&#34;:{&#34;type&#34;:&#34;browser&#34;,&#34;browserDetails&#34;:{&#34;viewportWidth&#34;:1280,&#34;viewportHeight&#34;:720}},&#34;placeContext&#34;:{&#34;localTimezoneOffset&#34;:420,&#34;localTime&#34;:&#34;2025-05-08T11:31:18.914-07:00&#34;},&#34;timestamp&#34;:&#34;2025-05-08T18:31:18.914Z&#34;,&#34;implementationDetails&#34;:{&#34;name&#34;:&#34;https://ns.adobe.com/experience/alloy/reactor&#34;,&#34;version&#34;:&#34;2.26.0&#43;2.29.0&#34;,&#34;environment&#34;:&#34;browser&#34;}},&#34;query&#34;:{&#34;personalization&#34;:{&#34;schemas&#34;:[&#34;https://ns.adobe.com/personalization/default-content-item&#34;,&#34;https://ns.adobe.com/personalization/html-content-item&#34;,&#34;https://ns.adobe.com/personalization/json-content-item&#34;,&#34;https://ns.adobe.com/personalization/redirect-item&#34;,&#34;https://ns.adobe.com/personalization/ruleset-item&#34;,&#34;https://ns.adobe.com/personalization/message/in-app&#34;,&#34;https://ns.adobe.com/personalization/message/content-card&#34;,&#34;https://ns.adobe.com/personalization/dom-action&#34;],&#34;decisionScopes&#34;:[&#34;__view__&#34;],&#34;surfaces&#34;:[&#34;web://www.crowdstrike.com/en-us&#34;]}}}],&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.googletagmanager.com/gtag/js?id=G-ZKTET1D58V", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/include/1746729300000/9d4udx6ceimp.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/scripts/bizible.js?account=crowdstrike.com", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCaf408a7d7bb64bba9c9ebd33e91577a5-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCbef2e7ebac224955b821d9e70110c78b-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC0c3a7a5da52745d29b8ee369060b69c6-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC0e765aea33c34bfe885facaf6fc7febb-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC30aa5696f5944a38a6fd12336395486c-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC702605e1c6554864bc71e30c9d81d484-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://munchkin.marketo.net/munchkin.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCe9011213d2114262a74c6c798d020dc2-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC9147058850554edca24cff1971dc8752-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCaf2d212b50ce48c1a6c3cd2be3fe5c2a-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCb6522d3edd5647029288904d3b3ea9e5-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/ipv?_biz_r=&amp;amp;_biz_h=-1719904874&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729079103&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=0&amp;amp;a=crowdstrike.com&amp;amp;rnd=480761&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729079104", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "image/avif,image/webp,image/apng,image/svg&#43;xml,image/*,*/*;q=0.8")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizibly.com/u?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729079104&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;a=crowdstrike.com&amp;amp;rnd=110587&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729079104", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "image/avif,image/webp,image/apng,image/svg&#43;xml,image/*,*/*;q=0.8")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/xdc.js?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_h=-1719904874&amp;amp;cdn_o=a&amp;amp;jsVer=4.25.04.18&amp;amp;a=crowdstrike.com", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC52d7b7dccc3548adaf8495e20b99dd8d-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ad.doubleclick.net/activity;u1=www.crowdstrike.com%2Fen-us%2F;cat=crowd0;src=12037336;type=crowd0?", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC429998db21a84d22962694934fe9c178-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC3f830817bbac469aa1a50e321edd3939-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC4428e44be4d744ddac8973a220b18817-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://connect.facebook.net/signals/config/950083805267950?v=2.9.201&amp;amp;r=stable&amp;amp;domain=www.crowdstrike.com&amp;amp;hme=9ebdfdd473ffce6bfe2267012c83f73483198ffe20d84139a2066b7682f827c0&amp;amp;ex_m=73%2C128%2C113%2C117%2C64%2C6%2C106%2C72%2C19%2C100%2C92%2C54%2C57%2C181%2C202%2C209%2C205%2C206%2C208%2C32%2C107%2C56%2C80%2C207%2C176%2C179%2C203%2C204%2C189%2C139%2C44%2C194%2C191%2C192%2C37%2C151%2C18%2C53%2C198%2C197%2C141%2C21%2C43%2C2%2C46%2C68%2C69%2C70%2C74%2C96%2C20%2C17%2C99%2C95%2C94%2C114%2C55%2C116%2C42%2C115%2C33%2C97%2C29%2C177%2C180%2C148%2C14%2C15%2C16%2C8%2C9%2C28%2C25%2C26%2C60%2C65%2C67%2C78%2C105%2C108%2C30%2C79%2C12%2C10%2C83%2C51%2C24%2C110%2C109%2C111%2C102%2C13%2C23%2C4%2C41%2C77%2C22%2C160%2C89%2C135%2C76%2C1%2C98%2C59%2C87%2C36%2C31%2C85%2C86%2C91%2C40%2C7%2C93%2C84%2C47%2C35%2C38%2C0%2C71%2C118%2C90%2C5%2C50%2C49%2C101%2C88%2C246%2C174%2C126%2C163%2C156%2C3%2C39%2C66%2C45%2C112%2C48%2C82%2C63%2C62%2C34%2C103%2C61%2C58%2C52%2C81%2C75%2C27%2C104%2C11%2C119", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://munchkin.marketo.net/164/munchkin.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ob.fishrobotflower.com/i/771439ae128c64ffe20e624628cb6c78.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://trk.techtarget.com/tracking.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://collector-20290.tvsquared.com/tv2track.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp-cdn.marketo.com/rtp-api/v1/rtp.js?aid=crowdstrike", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.redditstatic.com/ads/pixel.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://snap.licdn.com/li.lms-analytics/insight.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.influ2.com/tracker?clid=62c7557e-d1e3-40fb-93c4-d7c306706e53", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://arttrk.com/pixel/?ad_log=referer&amp;amp;action=lead&amp;amp;pixid=65acb722-b139-45a2-9a22-2e620e6d32b8", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "image/avif,image/webp,image/apng,image/svg&#43;xml,image/*,*/*;q=0.8")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ct.capterra.com/capterra_tracker.gif?vid=2104298&amp;amp;vkey=884c38bc6ebbb2426278e18b331d9004", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://edge.adobedc.net/ee/or2/v1/identity/acquire?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=23b5412d-0b01-4c08-b888-50320407803a", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true,&#34;entries&#34;:[{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_identity&#34;,&#34;value&#34;:&#34;CiYyODE4MDkyODQ4NjE4NzUyODY1MDM5MjUzNTg4NzAxMjI5MTU4OFISCJ2jqonrMhABGAEqA09SMjAA8AGdo6qJ6zI=&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_cluster&#34;,&#34;value&#34;:&#34;or2&#34;}]}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://edge.adobedc.net/ee/or2/v1/interact?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=d7735587-6597-4302-8f87-7b3545e81c37", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;events&#34;:[{&#34;xdm&#34;:{&#34;eventType&#34;:&#34;web.webpagedetails.pageViews&#34;,&#34;eventMergeId&#34;:&#34;75cfb0ab-4767-47ed-91e9-50bfe24f7ea1&#34;,&#34;web&#34;:{&#34;webPageDetails&#34;:{&#34;URL&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;name&#34;:&#34;/en-us/&#34;,&#34;pageViews&#34;:{&#34;value&#34;:1}},&#34;webReferrer&#34;:{&#34;URL&#34;:&#34;&#34;}},&#34;device&#34;:{&#34;screenHeight&#34;:720,&#34;screenWidth&#34;:1280,&#34;screenOrientation&#34;:&#34;landscape&#34;},&#34;environment&#34;:{&#34;type&#34;:&#34;browser&#34;,&#34;browserDetails&#34;:{&#34;viewportWidth&#34;:1280,&#34;viewportHeight&#34;:720}},&#34;placeContext&#34;:{&#34;localTimezoneOffset&#34;:420,&#34;localTime&#34;:&#34;2025-05-08T11:31:19.154-07:00&#34;},&#34;timestamp&#34;:&#34;2025-05-08T18:31:19.154Z&#34;,&#34;implementationDetails&#34;:{&#34;name&#34;:&#34;https://ns.adobe.com/experience/alloy/reactor&#34;,&#34;version&#34;:&#34;2.26.0&#43;2.29.0&#34;,&#34;environment&#34;:&#34;browser&#34;},&#34;_experience&#34;:{&#34;analytics&#34;:{&#34;event1to100&#34;:{&#34;event50&#34;:{&#34;id&#34;:&#34;Page Views&#34;,&#34;value&#34;:1}},&#34;customDimensions&#34;:{&#34;eVars&#34;:{&#34;eVar1&#34;:&#34;/en-us/&#34;,&#34;eVar2&#34;:&#34;www.crowdstrike.com&#34;,&#34;eVar3&#34;:&#34;&#34;,&#34;eVar4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;eVar6&#34;:&#34;dir&#34;,&#34;eVar7&#34;:&#34;&#34;,&#34;eVar10&#34;:&#34;35a249abd9e94512be1326ac06e4ddd5&#34;,&#34;eVar11&#34;:&#34;&#34;,&#34;eVar13&#34;:&#34;&#34;,&#34;eVar14&#34;:&#34;&#34;,&#34;eVar15&#34;:null,&#34;eVar16&#34;:&#34;bg41,c0001,c0002,c0003,c0004&#34;,&#34;eVar17&#34;:&#34;&#34;,&#34;eVar18&#34;:&#34;&#34;,&#34;eVar19&#34;:&#34;CrowdStrike-AEM|production|2025-05-05T15:24:36Z|Analytics : Global PageView | DOM Ready | Event-50&#34;,&#34;eVar20&#34;:&#34;/en-us/&#34;,&#34;eVar26&#34;:&#34;in&#34;,&#34;eVar50&#34;:&#34;cookie does not exist&#34;,&#34;eVar75&#34;:&#34;&#34;,&#34;eVar120&#34;:&#34;&#34;,&#34;eVar121&#34;:&#34;&#34;,&#34;eVar111&#34;:&#34;crowdstrike:homepage&#34;,&#34;eVar112&#34;:&#34;crowdstrike:en-us&#34;,&#34;eVar113&#34;:&#34;&#34;,&#34;eVar114&#34;:&#34;&#34;,&#34;eVar115&#34;:&#34;crowdstrike:en-us&#34;,&#34;eVar110&#34;:&#34;0.9205455244312417_1746729079155&#34;},&#34;props&#34;:{&#34;prop1&#34;:&#34;/en-us/&#34;,&#34;prop2&#34;:&#34;www.crowdstrike.com&#34;,&#34;prop4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;prop6&#34;:&#34;&#34;,&#34;prop10&#34;:&#34;crowdstrike:homepage&#34;,&#34;prop11&#34;:&#34;crowdstrike:en-us&#34;,&#34;prop12&#34;:&#34;&#34;,&#34;prop13&#34;:&#34;&#34;,&#34;prop14&#34;:&#34;crowdstrike:en-us&#34;}}}}},&#34;data&#34;:{&#34;__adobe&#34;:{&#34;target&#34;:{&#34;inTrial&#34;:&#34;false&#34;}}}}],&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true,&#34;entries&#34;:[{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_identity&#34;,&#34;value&#34;:&#34;CiYyODE4MDkyODQ4NjE4NzUyODY1MDM5MjUzNTg4NzAxMjI5MTU4OFISCJ2jqonrMhABGAEqA09SMjAA8AGdo6qJ6zI=&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_cluster&#34;,&#34;value&#34;:&#34;or2&#34;}]}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://edge.adobedc.net/ee/or2/v1/interact?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=7e6f42b5-795a-46e8-a906-fcafa82ba829", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;events&#34;:[{&#34;xdm&#34;:{&#34;_experience&#34;:{&#34;decisioning&#34;:{&#34;propositions&#34;:[{&#34;id&#34;:&#34;AT:eyJhY3Rpdml0eUlkIjoiMjc4NzM5IiwiZXhwZXJpZW5jZUlkIjoiMCJ9&#34;,&#34;scope&#34;:&#34;__view__&#34;,&#34;scopeDetails&#34;:{&#34;decisionProvider&#34;:&#34;TGT&#34;,&#34;activity&#34;:{&#34;id&#34;:&#34;278739&#34;},&#34;experience&#34;:{&#34;id&#34;:&#34;0&#34;},&#34;strategies&#34;:[{&#34;step&#34;:&#34;entry&#34;,&#34;trafficType&#34;:&#34;0&#34;},{&#34;step&#34;:&#34;display&#34;,&#34;trafficType&#34;:&#34;0&#34;}],&#34;characteristics&#34;:{&#34;eventToken&#34;:&#34;VFuU&#43;iZcTRWsijggTUoagJ3BNhkKdJh9FdggP2851ydw54XAlJ3xNRImNE7jBWgsgJd3DNmZFmAmR22YP7jXhw==&#34;},&#34;correlationID&#34;:&#34;278739:0:0&#34;}},{&#34;id&#34;:&#34;AT:eyJhY3Rpdml0eUlkIjoiMjcxNDMyIiwiZXhwZXJpZW5jZUlkIjoiMCJ9&#34;,&#34;scope&#34;:&#34;__view__&#34;,&#34;scopeDetails&#34;:{&#34;decisionProvider&#34;:&#34;TGT&#34;,&#34;activity&#34;:{&#34;id&#34;:&#34;271432&#34;},&#34;experience&#34;:{&#34;id&#34;:&#34;0&#34;},&#34;strategies&#34;:[{&#34;step&#34;:&#34;entry&#34;,&#34;trafficType&#34;:&#34;0&#34;},{&#34;step&#34;:&#34;display&#34;,&#34;trafficType&#34;:&#34;0&#34;}],&#34;characteristics&#34;:{&#34;eventToken&#34;:&#34;WF5pAfk2GLrxOsClXg3GXZ3BNhkKdJh9FdggP2851ydw54XAlJ3xNRImNE7jBWgsgJd3DNmZFmAmR22YP7jXhw==&#34;},&#34;correlationID&#34;:&#34;271432:0:0&#34;}},{&#34;id&#34;:&#34;AT:eyJhY3Rpdml0eUlkIjoiMjcxNTEzIiwiZXhwZXJpZW5jZUlkIjoiMCJ9&#34;,&#34;scope&#34;:&#34;__view__&#34;,&#34;scopeDetails&#34;:{&#34;decisionProvider&#34;:&#34;TGT&#34;,&#34;activity&#34;:{&#34;id&#34;:&#34;271513&#34;},&#34;experience&#34;:{&#34;id&#34;:&#34;0&#34;},&#34;strategies&#34;:[{&#34;step&#34;:&#34;entry&#34;,&#34;trafficType&#34;:&#34;0&#34;},{&#34;step&#34;:&#34;display&#34;,&#34;trafficType&#34;:&#34;0&#34;}],&#34;characteristics&#34;:{&#34;eventToken&#34;:&#34;8npzxVENtQ1VNij1pCFlxp3BNhkKdJh9FdggP2851ydw54XAlJ3xNRImNE7jBWgsgJd3DNmZFmAmR22YP7jXhw==&#34;},&#34;correlationID&#34;:&#34;271513:0:0&#34;}}],&#34;propositionEventType&#34;:{&#34;display&#34;:1}}},&#34;eventType&#34;:&#34;decisioning.propositionDisplay&#34;,&#34;web&#34;:{&#34;webPageDetails&#34;:{&#34;URL&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;},&#34;webReferrer&#34;:{&#34;URL&#34;:&#34;&#34;}},&#34;device&#34;:{&#34;screenHeight&#34;:720,&#34;screenWidth&#34;:1280,&#34;screenOrientation&#34;:&#34;landscape&#34;},&#34;environment&#34;:{&#34;type&#34;:&#34;browser&#34;,&#34;browserDetails&#34;:{&#34;viewportWidth&#34;:1280,&#34;viewportHeight&#34;:720}},&#34;placeContext&#34;:{&#34;localTimezoneOffset&#34;:420,&#34;localTime&#34;:&#34;2025-05-08T11:31:19.200-07:00&#34;},&#34;timestamp&#34;:&#34;2025-05-08T18:31:19.200Z&#34;,&#34;implementationDetails&#34;:{&#34;name&#34;:&#34;https://ns.adobe.com/experience/alloy/reactor&#34;,&#34;version&#34;:&#34;2.26.0&#43;2.29.0&#34;,&#34;environment&#34;:&#34;browser&#34;}}}],&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true,&#34;entries&#34;:[{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_identity&#34;,&#34;value&#34;:&#34;CiYyODE4MDkyODQ4NjE4NzUyODY1MDM5MjUzNTg4NzAxMjI5MTU4OFISCJ2jqonrMhABGAEqA09SMjAA8AGdo6qJ6zI=&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_cluster&#34;,&#34;value&#34;:&#34;or2&#34;}]}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC8a89dfd8b83c4511b4564c72e152a541-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://connect.facebook.net/signals/config/992980065451679?v=2.9.201&amp;amp;r=stable&amp;amp;domain=www.crowdstrike.com&amp;amp;hme=9ebdfdd473ffce6bfe2267012c83f73483198ffe20d84139a2066b7682f827c0&amp;amp;ex_m=73%2C128%2C113%2C117%2C64%2C6%2C106%2C72%2C19%2C100%2C92%2C54%2C57%2C181%2C202%2C209%2C205%2C206%2C208%2C32%2C107%2C56%2C80%2C207%2C176%2C179%2C203%2C204%2C189%2C139%2C44%2C194%2C191%2C192%2C37%2C151%2C18%2C53%2C198%2C197%2C141%2C21%2C43%2C2%2C46%2C68%2C69%2C70%2C74%2C96%2C20%2C17%2C99%2C95%2C94%2C114%2C55%2C116%2C42%2C115%2C33%2C97%2C29%2C177%2C180%2C148%2C14%2C15%2C16%2C8%2C9%2C28%2C25%2C26%2C60%2C65%2C67%2C78%2C105%2C108%2C30%2C79%2C12%2C10%2C83%2C51%2C24%2C110%2C109%2C111%2C102%2C13%2C23%2C4%2C41%2C77%2C22%2C160%2C89%2C135%2C76%2C1%2C98%2C59%2C87%2C36%2C31%2C85%2C86%2C91%2C40%2C7%2C93%2C84%2C47%2C35%2C38%2C0%2C71%2C118%2C90%2C5%2C50%2C49%2C101%2C88%2C246%2C174%2C126%2C163%2C156%2C3%2C39%2C66%2C45%2C112%2C48%2C82%2C63%2C62%2C34%2C103%2C61%2C58%2C52%2C81%2C75%2C27%2C104%2C11%2C119", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.facebook.com/tr/?id=950083805267950&amp;amp;ev=PageView&amp;amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;rl=&amp;amp;if=false&amp;amp;ts=1746729079201&amp;amp;sw=1280&amp;amp;sh=720&amp;amp;v=2.9.201&amp;amp;r=stable&amp;amp;a=adobe_launch&amp;amp;ec=0&amp;amp;o=28&amp;amp;it=1746729079168&amp;amp;coo=false&amp;amp;exp=k0&amp;amp;rqm=GET", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://281-obq-266.mktoresp.com/webevents/visitWebPage?_mchNc=1746729079204&amp;amp;_mchCn=&amp;amp;_mchId=281-OBQ-266&amp;amp;_mchTk=_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;_mchHo=www.crowdstrike.com&amp;amp;_mchPo=&amp;amp;_mchRu=%2Fen-us%2F&amp;amp;_mchPc=https%3A&amp;amp;_mchVr=164&amp;amp;aip=1&amp;amp;_mchEcid=06D71E9261F941560A495CD6%40AdobeOrg%3A%3A28180928486187528650392535887012291588&amp;amp;_mchHa=&amp;amp;_mchRe=&amp;amp;_mchQp=", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ad.doubleclick.net/activity;dc_pre=CPrKqfLAlI0DFeerjggdszgDLg;u1=www.crowdstrike.com%2Fen-us%2F;cat=crowd0;src=12037336;type=crowd0?", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCe54f772526294e1696f6be4e42252bd6-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ad.doubleclick.net/activity;u1=www.crowdstrike.com%2Fen-us%2F;cat=crowd00;src=12037336;type=crowd0?", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC13bb3acaad764c19a86cb66a819d9b01-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://t.contentsquare.net/uxa/184b355acd0d7.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://edge.adobedc.net/ee/or2/v1/interact?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=06dc78ed-dff6-4d81-a486-767e4dd14325", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;events&#34;:[{&#34;xdm&#34;:{&#34;eventType&#34;:&#34;web.webinteraction.linkClicks&#34;,&#34;eventMergeId&#34;:&#34;75cfb0ab-4767-47ed-91e9-50bfe24f7ea1&#34;,&#34;web&#34;:{&#34;webPageDetails&#34;:{&#34;URL&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;},&#34;webReferrer&#34;:{&#34;URL&#34;:&#34;&#34;},&#34;webInteraction&#34;:{&#34;URL&#34;:&#34;/en-us/&#34;,&#34;name&#34;:&#34;6sense&#34;,&#34;type&#34;:&#34;other&#34;,&#34;linkClicks&#34;:{&#34;value&#34;:1}}},&#34;device&#34;:{&#34;screenHeight&#34;:720,&#34;screenWidth&#34;:1280,&#34;screenOrientation&#34;:&#34;landscape&#34;},&#34;environment&#34;:{&#34;type&#34;:&#34;browser&#34;,&#34;browserDetails&#34;:{&#34;viewportWidth&#34;:1280,&#34;viewportHeight&#34;:720}},&#34;placeContext&#34;:{&#34;localTimezoneOffset&#34;:420,&#34;localTime&#34;:&#34;2025-05-08T11:31:19.474-07:00&#34;},&#34;timestamp&#34;:&#34;2025-05-08T18:31:19.474Z&#34;,&#34;implementationDetails&#34;:{&#34;name&#34;:&#34;https://ns.adobe.com/experience/alloy/reactor&#34;,&#34;version&#34;:&#34;2.26.0&#43;2.29.0&#34;,&#34;environment&#34;:&#34;browser&#34;},&#34;_experience&#34;:{&#34;analytics&#34;:{&#34;event1to100&#34;:{&#34;event53&#34;:{&#34;id&#34;:&#34;6sense Load&#34;,&#34;value&#34;:1}},&#34;customDimensions&#34;:{&#34;eVars&#34;:{&#34;eVar16&#34;:&#34;bg41,c0001,c0002,c0003,c0004&#34;,&#34;eVar19&#34;:&#34;CrowdStrike-AEM|production|2025-05-05T15:24:36Z|6Sense AA Integration&#34;,&#34;eVar26&#34;:&#34;in&#34;,&#34;eVar1&#34;:&#34;/en-us/&#34;,&#34;eVar2&#34;:&#34;www.crowdstrike.com&#34;,&#34;eVar3&#34;:&#34;&#34;,&#34;eVar4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;eVar6&#34;:&#34;dir&#34;,&#34;eVar7&#34;:&#34;&#34;,&#34;eVar15&#34;:null,&#34;eVar17&#34;:&#34;&#34;,&#34;eVar18&#34;:&#34;&#34;,&#34;eVar20&#34;:&#34;/en-us/&#34;,&#34;eVar111&#34;:&#34;crowdstrike:homepage&#34;,&#34;eVar112&#34;:&#34;crowdstrike:en-us&#34;,&#34;eVar113&#34;:&#34;&#34;,&#34;eVar114&#34;:&#34;&#34;,&#34;eVar115&#34;:&#34;crowdstrike:en-us&#34;},&#34;lists&#34;:{&#34;list2&#34;:{&#34;list&#34;:[{&#34;value&#34;:&#34;Domain : burtfeldman.com&#34;},{&#34;value&#34;:&#34;Name : Burt Feldman&#34;},{&#34;value&#34;:&#34;Region : Northern America&#34;},{&#34;value&#34;:&#34;Country : United States&#34;},{&#34;value&#34;:&#34;State : Arizona&#34;},{&#34;value&#34;:&#34;City : Scottsdale&#34;},{&#34;value&#34;:&#34;CompanyId : 249220946ae141e&#34;},{&#34;value&#34;:&#34;Country_iso_code : US&#34;},{&#34;value&#34;:&#34;Address : 7509 E 1ST ST&#34;},{&#34;value&#34;:&#34;Zip : 85251&#34;},{&#34;value&#34;:&#34;Phone : &#43;1 480-945-1800&#34;},{&#34;value&#34;:&#34;Employee_range : 0 - 9&#34;},{&#34;value&#34;:&#34;Revenue_range : $5M - $10M&#34;},{&#34;value&#34;:&#34;Employee_count : 1&#34;},{&#34;value&#34;:&#34;Annual_revenue : 5772000&#34;},{&#34;value&#34;:&#34;State_code : AZ&#34;},{&#34;value&#34;:&#34;GeoIP_country : United States&#34;},{&#34;value&#34;:&#34;GeoIP_state : California&#34;},{&#34;value&#34;:&#34;GeoIP_city : Sunnyvale&#34;},{&#34;value&#34;:&#34;Company_match : Match&#34;},{&#34;value&#34;:&#34;Additional_comment : Company name or domain match was found&#34;},{&#34;value&#34;:&#34;Sic : 8111&#34;},{&#34;value&#34;:&#34;Sic_description : Legal Services&#34;},{&#34;value&#34;:&#34;Naics : 541110&#34;},{&#34;value&#34;:&#34;Naics_description : Offices of Lawyers&#34;},{&#34;value&#34;:&#34;Industry 1 : Business Services&#34;},{&#34;value&#34;:&#34;Subindustry 1 : Legal&#34;},{&#34;value&#34;:&#34;product : crowdstrike&#34;},{&#34;value&#34;:&#34;intent_score : 0&#34;},{&#34;value&#34;:&#34;buying_stage : Target&#34;},{&#34;value&#34;:&#34;profile_score : 72&#34;},{&#34;value&#34;:&#34;profile_fit : Moderate&#34;},{&#34;value&#34;:&#34;is_6qa : false&#34;},{&#34;value&#34;:&#34;product_display_name : Crowdstrike&#34;},{&#34;value&#34;:&#34;Segment_name : TAM - Theatre - AMER&#34;},{&#34;value&#34;:&#34;Segment_id : 386105&#34;},{&#34;value&#34;:&#34;Segment_name : Brand Marketing | Olympics Targeting Segment | US |&#34;},{&#34;value&#34;:&#34;Segment_id : 723605&#34;},{&#34;value&#34;:&#34;Segment_name : GTM _ AMER _ SMB CORP&#34;},{&#34;value&#34;:&#34;Segment_id : 826118&#34;},{&#34;value&#34;:&#34;Segment_name : TAM - AMER - US&#34;},{&#34;value&#34;:&#34;Segment_id : 792640&#34;},{&#34;value&#34;:&#34;Segment_name : Brand Marketing | Strong/Moderate Target &#43; Awareness Stages (for Google)&#34;},{&#34;value&#34;:&#34;Segment_id : 716636&#34;}]}},&#34;props&#34;:{&#34;prop1&#34;:&#34;/en-us/&#34;,&#34;prop2&#34;:&#34;www.crowdstrike.com&#34;,&#34;prop4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;prop6&#34;:&#34;&#34;,&#34;prop10&#34;:&#34;crowdstrike:homepage&#34;,&#34;prop11&#34;:&#34;crowdstrike:en-us&#34;,&#34;prop12&#34;:&#34;&#34;,&#34;prop13&#34;:&#34;&#34;,&#34;prop14&#34;:&#34;crowdstrike:en-us&#34;}}}}},&#34;query&#34;:{&#34;personalization&#34;:{&#34;schemas&#34;:[&#34;https://ns.adobe.com/personalization/default-content-item&#34;,&#34;https://ns.adobe.com/personalization/html-content-item&#34;,&#34;https://ns.adobe.com/personalization/json-content-item&#34;,&#34;https://ns.adobe.com/personalization/redirect-item&#34;,&#34;https://ns.adobe.com/personalization/ruleset-item&#34;,&#34;https://ns.adobe.com/personalization/message/in-app&#34;,&#34;https://ns.adobe.com/personalization/message/content-card&#34;,&#34;https://ns.adobe.com/personalization/dom-action&#34;],&#34;decisionScopes&#34;:[&#34;__view__&#34;,&#34;mbox-sixth-sense&#34;],&#34;surfaces&#34;:[]}},&#34;data&#34;:{&#34;__adobe&#34;:{&#34;target&#34;:{&#34;type&#34;:&#34;6sense&#34;,&#34;6sense&#34;:&#34;true&#34;,&#34;profile.6sense_domain&#34;:&#34;burtfeldman.com&#34;,&#34;profile.6sense_name&#34;:&#34;Burt Feldman&#34;,&#34;profile.6sense_region&#34;:&#34;Northern America&#34;,&#34;profile.6sense_country&#34;:&#34;United States&#34;,&#34;profile.6sense_state&#34;:&#34;Arizona&#34;,&#34;profile.6sense_city&#34;:&#34;Scottsdale&#34;,&#34;profile.6sense_companyid&#34;:&#34;249220946ae141e&#34;,&#34;profile.6sense_country_iso_code&#34;:&#34;US&#34;,&#34;profile.6sense_address&#34;:&#34;7509 E 1ST ST&#34;,&#34;profile.6sense_zip&#34;:&#34;85251&#34;,&#34;profile.6sense_phone&#34;:&#34;&#43;1 480-945-1800&#34;,&#34;profile.6sense_employee_range&#34;:&#34;0 - 9&#34;,&#34;profile.6sense_revenue_range&#34;:&#34;$5M - $10M&#34;,&#34;profile.6sense_employee_count&#34;:&#34;1&#34;,&#34;profile.6sense_annual_revenue&#34;:&#34;5772000&#34;,&#34;profile.6sense_state_code&#34;:&#34;AZ&#34;,&#34;profile.6sense_geoip_country&#34;:&#34;United States&#34;,&#34;profile.6sense_geoip_state&#34;:&#34;California&#34;,&#34;profile.6sense_geoip_city&#34;:&#34;Sunnyvale&#34;,&#34;profile.6sense_company_match&#34;:&#34;Match&#34;,&#34;profile.6sense_additional_comment&#34;:&#34;Company name or domain match was found&#34;,&#34;profile.6sense_sic&#34;:&#34;8111&#34;,&#34;profile.6sense_sic_description&#34;:&#34;Legal Services&#34;,&#34;profile.6sense_naics&#34;:&#34;541110&#34;,&#34;profile.6sense_naics_description&#34;:&#34;Offices of Lawyers&#34;,&#34;profile.6sense_industry 1&#34;:&#34;Business Services&#34;,&#34;profile.6sense_subindustry 1&#34;:&#34;Legal&#34;,&#34;profile.6sense_product&#34;:&#34;crowdstrike&#34;,&#34;profile.6sense_intent_score&#34;:&#34;0&#34;,&#34;profile.6sense_buying_stage&#34;:&#34;Target&#34;,&#34;profile.6sense_profile_score&#34;:&#34;72&#34;,&#34;profile.6sense_profile_fit&#34;:&#34;Moderate&#34;,&#34;profile.6sense_is_6qa&#34;:&#34;false&#34;,&#34;profile.6sense_product_display_name&#34;:&#34;Crowdstrike&#34;,&#34;profile.6sense_segment_name&#34;:&#34;Brand Marketing | Strong/Moderate Target &#43; Awareness Stages (for Google)&#34;,&#34;profile.6sense_segment_id&#34;:&#34;716636&#34;}}}}],&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true,&#34;entries&#34;:[{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_identity&#34;,&#34;value&#34;:&#34;CiYyODE4MDkyODQ4NjE4NzUyODY1MDM5MjUzNTg4NzAxMjI5MTU4OFISCJ2jqonrMhABGAEqA09SMjAA8AGdo6qJ6zI=&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_cluster&#34;,&#34;value&#34;:&#34;or2&#34;}]}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC8a19ad85a7d044158f34540c69c7d8b4-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCca072a9a5f2f46d19d4d38bdcb9d6873-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://px.ads.linkedin.com/attribution_trigger?pid=64444&amp;amp;time=1746729079483&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://px.ads.linkedin.com/collect?v=2&amp;amp;fmt=js&amp;amp;pid=64444&amp;amp;time=1746729079483&amp;amp;li_adsId=f8924343-8fd8-45cc-9a32-420fa0932ae3&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://pixel-config.reddit.com/pixels/t2_2n40s6z5/config", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.redditstatic.com/ads/conversions-config/v1/pixel/config/t2_2n40s6z5_telemetry", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://alb.reddit.com/rp.gif?ts=1746729079487&amp;amp;id=t2_2n40s6z5&amp;amp;event=PageVisit&amp;amp;m.itemCount=&amp;amp;m.value=&amp;amp;m.valueDecimal=&amp;amp;m.currency=&amp;amp;m.transactionId=&amp;amp;m.customEventName=&amp;amp;m.products=&amp;amp;m.conversionId=&amp;amp;uuid=2bc44b75-14f8-4772-a1c1-5d39c30b5042&amp;amp;aaid=&amp;amp;em=&amp;amp;pn=&amp;amp;external_id=&amp;amp;idfa=&amp;amp;integration=reddit&amp;amp;partner=&amp;amp;opt_out=0&amp;amp;sh=1280&amp;amp;sw=720&amp;amp;v=rdt_00917691&amp;amp;dpm=&amp;amp;dpcc=&amp;amp;dprc=", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://t.influ2.com/u/?cb=1746729079487", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.google-analytics.com/g/collect?v=2&amp;amp;tid=G-ZKTET1D58V&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;_p=1746729079180&amp;amp;gcd=13l3l3l3l1l1&amp;amp;npa=0&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;gdid=dYWJhMj&amp;amp;cid=462250241.1746729079&amp;amp;ul=en-us&amp;amp;are=1&amp;amp;frm=0&amp;amp;pscdl=noapi&amp;amp;_geo=1&amp;amp;_rdi=1&amp;amp;_s=1&amp;amp;sid=1746729079&amp;amp;sct=1&amp;amp;seg=0&amp;amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;dt=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;en=page_view&amp;amp;_fv=1&amp;amp;_nsi=1&amp;amp;_ss=1&amp;amp;_ee=1&amp;amp;tfd=1360", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.google.com/ccm/collect?tid=AW-797629828&amp;amp;en=page_view&amp;amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;scrsrc=www.googletagmanager.com&amp;amp;frm=0&amp;amp;rnd=2104649365.1746729079&amp;amp;dt=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;auid=1475816545.1746729079&amp;amp;navt=n&amp;amp;npa=0&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;tft=1746729079497&amp;amp;tfd=1363&amp;amp;apve=1&amp;amp;apvf=sb", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://googleads.g.doubleclick.net/pagead/viewthroughconversion/797629828/?random=1746729079496&amp;amp;cv=11&amp;amp;fst=1746729079496&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;data=event%3Dgtag.config&amp;amp;rfmt=3&amp;amp;fmt=4", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://td.doubleclick.net/td/rul/797629828?random=1746729079496&amp;amp;cv=11&amp;amp;fst=1746729079496&amp;amp;fmt=3&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;data=event%3Dgtag.config", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://googleads.g.doubleclick.net/pagead/viewthroughconversion/952416460/?random=1746729079503&amp;amp;cv=11&amp;amp;fst=1746729079503&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;data=event%3Dgtag.config&amp;amp;rfmt=3&amp;amp;fmt=4", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://td.doubleclick.net/td/rul/952416460?random=1746729079503&amp;amp;cv=11&amp;amp;fst=1746729079503&amp;amp;fmt=3&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;data=event%3Dgtag.config", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.datadoghq-browser-agent.com/us1/v5/datadog-rum.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://rtp-static.marketo.com/rtp/libs/jquery/3.7.0/jquery.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://rtp-static.marketo.com/rtp/libs/jquery-ui-insightera-custom-1.9.6.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "text/css,*/*;q=0.1")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/trw?aid=crowdstrike&amp;amp;trwv.uid=crowdstrike-1746729079514-1f81ceec&amp;amp;trwv.vc=1&amp;amp;trwsa.sid=crowdstrike-1746729079515-66d02f45&amp;amp;trwsb.cpv=1&amp;amp;ctzo=-07:00&amp;amp;uri=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;pm=&amp;amp;viewedTypes=&amp;amp;rts=1746729079515", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://rtp-static.marketo.com/rtp/libs/ga-integration-2.0.5.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ad.doubleclick.net/activity;dc_pre=COf3uPLAlI0DFYbqrQAdK0ceQQ;u1=www.crowdstrike.com%2Fen-us%2F;cat=crowd00;src=12037336;type=crowd0?", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://adservice.google.com/ddm/fls/z/dc_pre=CPrKqfLAlI0DFeerjggdszgDLg;u1=www.crowdstrike.com%2Fen-us%2F;cat=crowd0;src=12037336;type=crowd0", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ibc-flow.techtarget.com/a/gif.gif?actTypeId=31&amp;amp;cid=3218843&amp;amp;r=1746729079517&amp;amp;ref=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;version=2.4", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != -1 {
		t.Errorf("Expected status %d, got %d", -1, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.facebook.com/tr/?id=992980065451679&amp;amp;ev=PageView&amp;amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;rl=&amp;amp;if=false&amp;amp;ts=1746729079523&amp;amp;sw=1280&amp;amp;sh=720&amp;amp;v=2.9.201&amp;amp;r=stable&amp;amp;ec=0&amp;amp;o=4126&amp;amp;fbp=fb.1.1746729079520.75795469452619173&amp;amp;ler=empty&amp;amp;it=1746729079168&amp;amp;coo=false&amp;amp;exp=k2&amp;amp;rqm=GET", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.googletagmanager.com/static/service_worker/5570/sw_iframe.html?origin=https%3A%2F%2Fwww.crowdstrike.com", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ibc-flow.techtarget.com/a/gif.gif?actTypeId=31&amp;amp;cid=3218843&amp;amp;r=1746729079530&amp;amp;ref=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;version=2.4", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != -1 {
		t.Errorf("Expected status %d, got %d", -1, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://adservice.google.com/ddm/fls/z/dc_pre=COf3uPLAlI0DFYbqrQAdK0ceQQ;u1=www.crowdstrike.com%2Fen-us%2F;cat=crowd00;src=12037336;type=crowd0", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/widget.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://obs.fishrobotflower.com/ct?id=42110&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;sf=0&amp;amp;tpi=&amp;amp;ch=cheq4ppc&amp;amp;uvid=undefined&amp;amp;tsf=0&amp;amp;tsfmi=&amp;amp;tsfu=&amp;amp;cb=1746729079646&amp;amp;hl=2&amp;amp;op=0&amp;amp;ag=1932998597&amp;amp;rand=8452781980417220858212893201002921909569170169186085261562527102269915086580215220293222&amp;amp;fs=1280x720&amp;amp;fst=1280x720&amp;amp;np=macintel&amp;amp;nv=google%20inc.&amp;amp;ref=&amp;amp;ss=1280x720&amp;amp;nc=0&amp;amp;at=&amp;amp;di=W1siZWYiLDQwNTVdLFsiYWJuY2giLDVdLFstMTMsIi0iXSxbLTIyLCJbXCJuXCIsXCJuXCJdIl0sWy0zOCwiaSwtMSwtMSwwLDAsMCwwLDIsNTYsMTAxLC0xLDAsNTE2LDUxNiwxMzk1LDEzOTYiXSxbLTMsIltdIl0sWy0xMCwiLSJdLFstNDIsIjg4MzM5OTAxNiJdLFstNDUsIjYyMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsNjMwLDU3MSw0OTMsNjA4LDc5Miw2NzcsMCwwLDAsMCwwIl0sWzEyLCJ7XCJjdHhcIjpcIndlYmdsXCIsXCJ2XCI6XCJnb29nbGUgaW5jLiAoYXBwbGUpXCIsXCJyXCI6XCJhbmdsZSAoYXBwbGUsIGFwcGxlIG00IHBybywgb3BlbmdsIDQuMSlcIixcInNsdlwiOlwid2ViZ2wgZ2xzbCBlcyAxLjAgKG9wZW5nbCBlcyBnbHNsIGVzIDEuMCBjaHJvbWl1bSlcIixcImd2ZXJcIjpcIndlYmdsIDEuMCAob3BlbmdsIGVzIDIuMCBjaHJvbWl1bSlcIixcImd2ZW5cIjpcIndlYmtpdFwiLFwiYmVuXCI6MixcIndnbFwiOjEsXCJncmVuXCI6XCJ3ZWJraXQgd2ViZ2xcIixcInNlZlwiOjI5NzA0MTcwOCxcInNlY1wiOlwiXCJ9Il0sWy0xOSwiWzAsMCwwLDAsMCwwLDEsMjQsMjQsXCItXCIsMTI4MCw3MjAsMTI4MCw3MjAsMTI4MCw3MjAsMTI4MCw3MjAsMCwwLDAsMCxcIi1cIixcIi1cIiwxMjgwLDcyMCwwXSJdLFstMjYsIntcInRqaHNcIjozNTEwMDAwMCxcInVqaHNcIjoyNjAwMDAwMCxcImpoc2xcIjozNzYwMDAwMDAwfSJdLFstMjcsIlswLDEwLDAsXCI0Z1wiLG51bGxdIl0sWy0zNywiLSJdLFstNTUsIjEiXSxbLTUyLCItIl0sWy02MSwie1wid2dzbFwiOlwiNDtwYWNrZWRfNHg4X2ludGVnZXJfZG90X3Byb2R1Y3Q7dW5yZXN0cmljdGVkX3BvaW50ZXJfcGFyYW1ldGVycztwb2ludGVyX2NvbXBvc2l0ZV9hY2Nlc3M7cmVhZG9ubHlfYW5kX3JlYWR3cml0ZV9zdG9yYWdlX3RleHR1cmVzO1wiLFwicGNmXCI6XCJiZ3JhOHVub3JtXCJ9Il0sWy02NCwiWzAsXCJXaW5kb3dzXCIsW3tcImJcIjpcIkNocm9taXVtXCIsXCJ2XCI6XCIxMzZcIn0se1wiYlwiOlwiSGVhZGxlc3NDaHJvbWVcIixcInZcIjpcIjEzNlwifSx7XCJiXCI6XCJOb3QuQS9CcmFuZFwiLFwidlwiOlwiOTlcIn1dXSJdLFstMTQsIi0iXSxbLTIxLCItIl0sWy0yMywiKyJdLFstMjQsIltcInNheXN3aG9cIiwwLFwiQ2hyb21lIDEzNlwiLDEsMV0iXSxbLTQwLCIzMyJdLFstNTYsImxhbmRzY2FwZS1wcmltYXJ5Il0sWy01LCItIl0sWy02LCJ7XCJ3XCI6W1wiMFwiLFwiT25lVHJ1c3RTdHViXCIsXCJPcHRhbm9uV3JhcHBlclwiLFwiJFwiLFwialF1ZXJ5XCIsXCJtYXRjaGVkXCIsXCJicm93c2VyXCIsXCJHcmFuaXRlXCIsXCJFbmxpZ2h0ZXJKU1wiLFwidmlkeWFyZEVtYmVkXCIsXCJzZXRJbW1lZGlhdGVcIixcImNsZWFySW1tZWRpYXRlXCIsXCJWaWR5YXJkVjRcIixcIlZpZHlhcmRcIixcIk90VHJ1c3RlZFR5cGVcIixcIl9fU1ZHX1NQUklURV9fXCIsXCJBZGRTZWFyY2hDbGllbnRcIixcIkFkZFNlYXJjaFVJXCIsXCJnc2FwVmVyc2lvbnNcIixcIkNRXCIsXCJfc2xpY2VkVG9BcnJheVwiLFwiX25vbkl0ZXJhYmxlUmVzdFwiLFwiX2l0ZXJhYmxlVG9BcnJheUxpbWl0XCIsXCJfYXJyYXlXaXRoSG9sZXNcIixcIl9jcmVhdGVGb3JPZkl0ZXJhdG9ySGVscGVyXCIsXCJfdW5zdXBwb3J0ZWRJdGVyYWJsZVRvQXJyYXlcIixcIl9hcnJheUxpa2VUb0FycmF5XCIsXCJfdHlwZW9mXCIsXCJDTVBcIixcImFkb2JlRGF0YUxheWVyXCIsXCJyZWFjdGl2ZUVsZW1lbnRWZXJzaW9uc1wiLFwibGl0SHRtbFZlcnNpb25zXCIsXCJsaXRFbGVtZW50VmVyc2lvbnNcIixcIkNTU1J1bGVQbHVnaW5cIixcIkN1c3RvbUVhc2VcIixcIkRyYXdTVkdQbHVnaW5cIixcIkVhc2VsUGx1Z2luXCIsXCJFYXNlUGFja1wiLFwiRXhwb1NjYWxlRWFzZVwiLFwiUm91Z2hFYXNlXCIsXCJTbG93TW9cIixcIkxpbmVhclwiLFwiUG93ZXIwXCIsXCJRdWFkXCIsXCJQb3dlcjFcIixcIkN1YmljXCIsXCJQb3dlcjJcIixcIlF1YXJ0XCIsXCJQb3dlcjNcIixcIlF1aW50XCJdLFwiblwiOltcInNheXN3aG9cIl0sXCJkXCI6W119Il0sWy05LCIrIl0sWy0xNywiMTIiXSxbLTIwLCI0NjIyNTAyNDEuMTc0NjcyOTA3OSJdLFstMjUsIi0iXSxbLTM0LCItIl0sWy0zNiwiW1wiMTYvOVwiLFwiMTYvOVwiXSJdLFstNDMsIjAwMDAwMDAxMDAwMDAwMDAxMDExMTAxMTAwMTExMTAxMDAwMDAxMCJdLFstNDksIi0iXSxbLTUzLCIwMDEiXSxbLTU4LCItIl0sWy02MCwiLSJdLFstNjIsIjgwIl0sWy0xLCItIl0sWy03LCItIl0sWy0xMSwie1widFwiOlwiXCIsXCJtXCI6W1wiZGVzY3JpcHRpb25cIixcIm9nOnRpdGxlXCIsXCJvZzpkZXNjcmlwdGlvblwiLFwidHdpdHRlcjp0aXRsZVwiLFwidHdpdHRlcjpkZXNjcmlwdGlvblwiXX0iXSxbLTMwLCJbXCJ2XCIsMF0iXSxbLTMzLCItIl0sWy03MSwiYTAxMDAxMDExMDAxMDAxMDEwMDAxMDEwMDExMTExMDEwMDAwMTAiXSxbLTQsIi0iXSxbLTgsIi0iXSxbLTEyLCJudWxsIl0sWy0xNiwiMCJdLFstMTgsIlswLDAsMCwxXSJdLFstMzIsIjAiXSxbLTUwLCItIl0sWy01NCwie1wiaFwiOltcIl8zXCIsXCIzMjk5OTEzNjlcIixcIjY4NDQwMTg5NlwiLFwiMzI5OTkxMzY5XCIsXCIxMjM4NDExODczXCIsXCIxMTc0OTg5NTU5XCIsXCJfMlwiLFwiMjM0MzM1MjMxNVwiXSxcImRcIjpbXSxcImJcIjpbXCJfMVwiLFwiMjA2Njc4MTQxMFwiLFwiMzk2NTA1MDMzXCIsXCI3OTUzOTU3MDlcIixcIjExOTYyNzcwNTVcIixcIjkxODc4Mzk2OVwiLFwiMTY5ODUyMDA3OVwiLFwiNDIwNjA3OTA3M1wiLFwiMzYyNzIyOTcxM1wiLFwiMTE3NDk4OTU1OVwiXSxcInNcIjoxfSJdLFstNTksImRlbmllZCJdLFstNjMsIjAiXSxbLTY1LCItIl0sWy0xNSwiLSJdLFstMjgsImVuLVVTIl0sWy0zOSwiW1wiMjAwMzAxMDdcIiwwLFwiR2Vja29cIixcIk5ldHNjYXBlXCIsXCJNb3ppbGxhXCIsbnVsbCxudWxsLHRydWUsOCxmYWxzZSxudWxsLDAsZmFsc2UsdHJ1ZSxudWxsLDAsdHJ1ZSx0cnVlXSJdLFstNDQsIjAsMCwwLDUiXSxbLTcyLCJFeFU9Il0sWy0yLCIzMixlWW5XV1o5KytudnU5cE0zT21wZmVFTUNtRWFxUUVBaUYwcENqWVZwcm9XbFpkbEJVRTE0QUZsVlVSY1pXT2lpQmxzYUVJVWdPa1VCSklRbnFkdEVreU01bHl5bHVlOXUzOW5qIl0sWy0yOSwiLSJdLFstMzEsInRydWUiXSxbLTM1LCJbMTc0NjcyOTA3OTY0NCw3XSJdLFstNDEsIi0iXSxbLTQ2LCIwIl0sWy01MSwiLSJdLFstNTcsIldFMFpYVlphVEZSY1YwMFhXa3RjV0UxY2ZGVmNWRnhYVFJrUlVVMU5TVW9ERmhaYVhWY1hXbFpXVWxCY1ZWaE9GMVpMWGhaYVZsZEtYRmRORmx0Y1hBZ01XdzVhRkZzUENnc1VEUXdKWEJRQUNRa0tGQUJhQVZzUENWc0tXd0FPQVJaMlRYaE1UVlo3VlZaYVVoZFRTZ01BQXcwUEFSQVZXRTBaZUV0TFdFQVhYVndaRVZGTlRVbEtBeFlXVmxzWFgxQktVVXRXVzFaTlgxVldUbHhMRjFwV1ZCWlFGZzRPQ0EwS0FGaGNDQXNCV2c4TlgxOWNDd2xjRHdzTkR3c0JXbHNQV2c0QkYxTktBd2dEREE0UENnMFFGVmhOR1VzWkVWRk5UVWxLQXhZV1Zsc1hYMUJLVVV0V1cxWk5YMVZXVGx4TEYxcFdWQlpRRmc0T0NBMEtBQT09Il0sWy02NiwiZ2VvbG9jYXRpb24sY2h1YWZ1bGx2ZXJzaW9ubGlzdCxjcm9zc29yaWdpbmlzb2xhdGVkLHNjcmVlbndha2Vsb2NrLHB1YmxpY2tleWNyZWRlbnRpYWxzZ2V0LHNoYXJlZHN0b3JhZ2VzZWxlY3R1cmwsY2h1YWFyY2gsYmx1ZXRvb3RoLGNvbXB1dGVwcmVzc3VyZSxjaHByZWZlcnNyZWR1Y2VkdHJhbnNwYXJlbmN5LGRlZmVycmVkZmV0Y2gsdXNiLGNoc2F2ZWRhdGEscHVibGlja2V5Y3JlZGVudGlhbHNjcmVhdGUsc2hhcmVkc3RvcmFnZSxkZWZlcnJlZGZldGNobWluaW1hbCxydW5hZGF1Y3Rpb24sY2hkb3dubGluayxjaHVhZm9ybWZhY3RvcnMsb3RwY3JlZGVudGlhbHMscGF5bWVudCxjaHVhLGNodWFtb2RlbCxjaGVjdCxhdXRvcGxheSxjYW1lcmEscHJpdmF0ZXN0YXRldG9rZW5pc3N1YW5jZSxhY2NlbGVyb21ldGVyLGNodWFwbGF0Zm9ybXZlcnNpb24saWRsZWRldGVjdGlvbixwcml2YXRlYWdncmVnYXRpb24saW50ZXJlc3Rjb2hvcnQsY2h2aWV3cG9ydGhlaWdodCxjYXB0dXJlZHN1cmZhY2Vjb250cm9sLGxvY2FsZm9udHMsY2h1YXBsYXRmb3JtLG1pZGksY2h1YWZ1bGx2ZXJzaW9uLHhyc3BhdGlhbHRyYWNraW5nLGNsaXBib2FyZHJlYWQsZ2FtZXBhZCxkaXNwbGF5Y2FwdHVyZSxrZXlib2FyZG1hcCxqb2luYWRpbnRlcmVzdGdyb3VwLGNod2lkdGgsY2hwcmVmZXJzcmVkdWNlZG1vdGlvbixicm93c2luZ3RvcGljcyxlbmNyeXB0ZWRtZWRpYSxneXJvc2NvcGUsc2VyaWFsLGNocnR0LGNodWFtb2JpbGUsd2luZG93bWFuYWdlbWVudCx1bmxvYWQsY2hkcHIsY2hwcmVmZXJzY29sb3JzY2hlbWUsY2h1YXdvdzY0LGF0dHJpYnV0aW9ucmVwb3J0aW5nLGZ1bGxzY3JlZW4saWRlbnRpdHljcmVkZW50aWFsc2dldCxwcml2YXRlc3RhdGV0b2tlbnJlZGVtcHRpb24saGlkLGNodWFiaXRuZXNzLHN0b3JhZ2VhY2Nlc3Msc3luY3hocixjaGRldmljZW1lbW9yeSxjaHZpZXdwb3J0d2lkdGgscGljdHVyZWlucGljdHVyZSxtYWduZXRvbWV0ZXIsY2xpcGJvYXJkd3JpdGUsbWljcm9waG9uZSJdLFstNjcsIi0iXSxbLTY4LCJbMjg2NjAzNTQ0OSxmdW5jdGlvbigpe2Zvcih2YXIgYSxjLGI9W10sZT0wO2U8YXJndW1lbnRzLmxlbmd0aDtlKyspYltlXT1hcmd1bWVudHNbZV07cmV0dXJuXCJzY3JpcHRcIj09PWJbMF0udG9Mb3dlckNhc2UoKXx8LTEhPT10LmluZGV4T2YoYlswXS50b0xvd2VyQ2FzZSgpKT8oYT1CLmJpbmQoZG9jdW1lbnQpLmFwcGx5KHZvaWQgMCxiKSxjPWEuc2V0QXR0cmlidXRlLmJpbmQoYSksT2JqZWN0LmRlZmluZVByb3BlcnRpZXMoYSx7c3JjOnsiXSxbLTY5LCJNYWNJbnRlbHxHb29nbGUgSW5jLnw4fDEyfFdpbmRvd3N8MCJdLFsiYm5jaCIsMTM4XSxbLTQ3LCJBbWVyaWNhL0xvc19BbmdlbGVzLGVuLVVTLGxhdG4sZ3JlZ29yeSJdLFstNDgsIltcIi1cIixcIi1cIixcIi1cIl0iXSxbLTcwLCItIl0sWyJkZGIiLCIwLDMyLDAsMCwxLDAsMCwwLDAsMCwwLDAsMCwwLDEsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwxLDQsMCwxLDAsMCwwLDAsMCwwLDAsMCwwLDAsMSw2MiwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMSwwIl0sWyJjYiIsIjAsMCwwLDAsMCwwLDAsMSwwLDAsMCwwLDMsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMiwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDEiXV0%3D&amp;amp;dep=0&amp;amp;pre=0&amp;amp;sdd=&amp;amp;cri=jfECv66HI9&amp;amp;pto=1512&amp;amp;ver=65&amp;amp;gac=462250241.1746729079&amp;amp;mei=&amp;amp;ap=&amp;amp;fe=1&amp;amp;duid=1.1746729079.FC7U58A4vZI5yYK0&amp;amp;suid=1.1746729079.KGNJ6zRpGQhCdArz&amp;amp;tuid=1.1746729079.hJyS90a8DamtJXqt&amp;amp;fbc=1.1746729079520.75795469452619173&amp;amp;gtm=WyJPbmVUcnVzdExvYWRlZCIsIk9wdGFub25Mb2FkZWQiLCJPbmVUcnVzdEdyb3Vwc1VwZGF0ZWQiXQ%3D%3D&amp;amp;it=146%2C1052%2C109&amp;amp;fbcl=-&amp;amp;gacl=-&amp;amp;gacsd=-&amp;amp;rtic=-&amp;amp;rtict=-&amp;amp;bgc=-&amp;amp;spa=1&amp;amp;urid=0&amp;amp;ab=jx.2.0%3B&amp;amp;sck=-&amp;amp;io=aGA2Oi17bmY2Og%3D%3D", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://edge.adobedc.net/ee/or2/v1/interact?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=f885747c-7369-4719-9398-9a8c0b7dba22", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;events&#34;:[{&#34;xdm&#34;:{&#34;eventType&#34;:&#34;web.webinteraction.linkClicks&#34;,&#34;eventMergeId&#34;:&#34;75cfb0ab-4767-47ed-91e9-50bfe24f7ea1&#34;,&#34;web&#34;:{&#34;webPageDetails&#34;:{&#34;URL&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;},&#34;webReferrer&#34;:{&#34;URL&#34;:&#34;&#34;},&#34;webInteraction&#34;:{&#34;URL&#34;:&#34;/en-us/&#34;,&#34;name&#34;:&#34;Target Data&#34;,&#34;type&#34;:&#34;other&#34;,&#34;linkClicks&#34;:{&#34;value&#34;:1}}},&#34;device&#34;:{&#34;screenHeight&#34;:720,&#34;screenWidth&#34;:1280,&#34;screenOrientation&#34;:&#34;landscape&#34;},&#34;environment&#34;:{&#34;type&#34;:&#34;browser&#34;,&#34;browserDetails&#34;:{&#34;viewportWidth&#34;:1280,&#34;viewportHeight&#34;:720}},&#34;placeContext&#34;:{&#34;localTimezoneOffset&#34;:420,&#34;localTime&#34;:&#34;2025-05-08T11:31:19.650-07:00&#34;},&#34;timestamp&#34;:&#34;2025-05-08T18:31:19.650Z&#34;,&#34;implementationDetails&#34;:{&#34;name&#34;:&#34;https://ns.adobe.com/experience/alloy/reactor&#34;,&#34;version&#34;:&#34;2.26.0&#43;2.29.0&#34;,&#34;environment&#34;:&#34;browser&#34;},&#34;_experience&#34;:{&#34;analytics&#34;:{&#34;event1to100&#34;:{&#34;event52&#34;:{&#34;id&#34;:&#34;Target Load&#34;,&#34;value&#34;:1}},&#34;customDimensions&#34;:{&#34;eVars&#34;:{&#34;eVar1&#34;:&#34;/en-us/&#34;,&#34;eVar2&#34;:&#34;www.crowdstrike.com&#34;,&#34;eVar3&#34;:&#34;&#34;,&#34;eVar4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;eVar6&#34;:&#34;dir&#34;,&#34;eVar7&#34;:&#34;&#34;,&#34;eVar10&#34;:&#34;35a249abd9e94512be1326ac06e4ddd5&#34;,&#34;eVar11&#34;:&#34;id:281-OBQ-266&amp;token:_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&#34;,&#34;eVar13&#34;:&#34;fb.1.1746729079520.75795469452619173&#34;,&#34;eVar14&#34;:&#34;GA1.1.462250241.1746729079&#34;,&#34;eVar15&#34;:&#34;28180928486187528650392535887012291588&#34;,&#34;eVar16&#34;:&#34;bg41,c0001,c0002,c0003,c0004&#34;,&#34;eVar17&#34;:&#34;&#34;,&#34;eVar18&#34;:&#34;&#34;,&#34;eVar19&#34;:&#34;CrowdStrike-AEM|production|2025-05-05T15:24:36Z|Adobe Target Data Collection&#34;,&#34;eVar101&#34;:&#34;Homepage FIXES,fix-wst-340-1,fix-wst-340-2&#34;,&#34;eVar102&#34;:&#34;Experience A,Experience A,Experience A&#34;,&#34;eVar103&#34;:&#34;Homepage FIXES:Experience A,fix-wst-340-1:Experience A,fix-wst-340-2:Experience A&#34;,&#34;eVar104&#34;:&#34;Homepage FIXES,fix-wst-340-1,fix-wst-340-2&#34;,&#34;eVar105&#34;:&#34;278739:0,271432:0,271513:0&#34;,&#34;eVar106&#34;:&#34;278739:0,271432:0,271513:0&#34;,&#34;eVar20&#34;:&#34;/en-us/&#34;,&#34;eVar26&#34;:&#34;in&#34;,&#34;eVar111&#34;:&#34;crowdstrike:homepage&#34;,&#34;eVar112&#34;:&#34;crowdstrike:en-us&#34;,&#34;eVar113&#34;:&#34;&#34;,&#34;eVar114&#34;:&#34;&#34;,&#34;eVar115&#34;:&#34;crowdstrike:en-us&#34;},&#34;listProps&#34;:{&#34;prop70&#34;:{&#34;values&#34;:&#34;&#34;},&#34;prop71&#34;:{&#34;values&#34;:&#34;&#34;},&#34;prop72&#34;:{&#34;values&#34;:&#34;&#34;},&#34;prop73&#34;:{&#34;values&#34;:&#34;Homepage FIXES,fix-wst-340-1,fix-wst-340-2&#34;},&#34;prop74&#34;:{&#34;values&#34;:&#34;Experience A,Experience A,Experience A&#34;},&#34;prop75&#34;:{&#34;values&#34;:&#34;Homepage FIXES:Experience A,fix-wst-340-1:Experience A,fix-wst-340-2:Experience A&#34;}},&#34;props&#34;:{&#34;prop1&#34;:&#34;/en-us/&#34;,&#34;prop2&#34;:&#34;www.crowdstrike.com&#34;,&#34;prop4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;prop6&#34;:&#34;&#34;,&#34;prop10&#34;:&#34;crowdstrike:homepage&#34;,&#34;prop11&#34;:&#34;crowdstrike:en-us&#34;,&#34;prop12&#34;:&#34;&#34;,&#34;prop13&#34;:&#34;&#34;,&#34;prop14&#34;:&#34;crowdstrike:en-us&#34;}}}}}}],&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true,&#34;entries&#34;:[{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_identity&#34;,&#34;value&#34;:&#34;CiYyODE4MDkyODQ4NjE4NzUyODY1MDM5MjUzNTg4NzAxMjI5MTU4OFISCJ2jqonrMhABGAEqA09SMjAA8AGdo6qJ6zI=&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_cluster&#34;,&#34;value&#34;:&#34;or2&#34;}]}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://px.ads.linkedin.com/collect?v=2&amp;amp;fmt=js&amp;amp;pid=64444&amp;amp;time=1746729079483&amp;amp;li_adsId=f8924343-8fd8-45cc-9a32-420fa0932ae3&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;cookiesTest=true", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/msg?a=2&amp;amp;sid=crowdstrike-1746729079515-66d02f45&amp;amp;aid=crowdstrike&amp;amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;viewedTypes=&amp;amp;0.08197428152563013&amp;amp;rts=1746729079668", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ext.chtbl.com/trackable.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://edge.adobedc.net/ee/or2/v1/interact?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=2495863c-3ad6-40d1-a0db-a26631d80024", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;events&#34;:[{&#34;xdm&#34;:{&#34;_experience&#34;:{&#34;decisioning&#34;:{&#34;propositions&#34;:[{&#34;id&#34;:&#34;AT:eyJhY3Rpdml0eUlkIjoiMjc4NzM5IiwiZXhwZXJpZW5jZUlkIjoiMCJ9&#34;,&#34;scope&#34;:&#34;__view__&#34;,&#34;scopeDetails&#34;:{&#34;decisionProvider&#34;:&#34;TGT&#34;,&#34;activity&#34;:{&#34;id&#34;:&#34;278739&#34;},&#34;experience&#34;:{&#34;id&#34;:&#34;0&#34;},&#34;strategies&#34;:[{&#34;step&#34;:&#34;entry&#34;,&#34;trafficType&#34;:&#34;0&#34;},{&#34;step&#34;:&#34;display&#34;,&#34;trafficType&#34;:&#34;0&#34;}],&#34;characteristics&#34;:{&#34;eventToken&#34;:&#34;VFuU&#43;iZcTRWsijggTUoagJ3BNhkKdJh9FdggP2851ydw54XAlJ3xNRImNE7jBWgsgJd3DNmZFmAmR22YP7jXhw==&#34;},&#34;correlationID&#34;:&#34;278739:0:0&#34;}},{&#34;id&#34;:&#34;AT:eyJhY3Rpdml0eUlkIjoiMjcxNDMyIiwiZXhwZXJpZW5jZUlkIjoiMCJ9&#34;,&#34;scope&#34;:&#34;__view__&#34;,&#34;scopeDetails&#34;:{&#34;decisionProvider&#34;:&#34;TGT&#34;,&#34;activity&#34;:{&#34;id&#34;:&#34;271432&#34;},&#34;experience&#34;:{&#34;id&#34;:&#34;0&#34;},&#34;strategies&#34;:[{&#34;step&#34;:&#34;entry&#34;,&#34;trafficType&#34;:&#34;0&#34;},{&#34;step&#34;:&#34;display&#34;,&#34;trafficType&#34;:&#34;0&#34;}],&#34;characteristics&#34;:{&#34;eventToken&#34;:&#34;WF5pAfk2GLrxOsClXg3GXZ3BNhkKdJh9FdggP2851ydw54XAlJ3xNRImNE7jBWgsgJd3DNmZFmAmR22YP7jXhw==&#34;},&#34;correlationID&#34;:&#34;271432:0:0&#34;}},{&#34;id&#34;:&#34;AT:eyJhY3Rpdml0eUlkIjoiMjcxNTEzIiwiZXhwZXJpZW5jZUlkIjoiMCJ9&#34;,&#34;scope&#34;:&#34;__view__&#34;,&#34;scopeDetails&#34;:{&#34;decisionProvider&#34;:&#34;TGT&#34;,&#34;activity&#34;:{&#34;id&#34;:&#34;271513&#34;},&#34;experience&#34;:{&#34;id&#34;:&#34;0&#34;},&#34;strategies&#34;:[{&#34;step&#34;:&#34;entry&#34;,&#34;trafficType&#34;:&#34;0&#34;},{&#34;step&#34;:&#34;display&#34;,&#34;trafficType&#34;:&#34;0&#34;}],&#34;characteristics&#34;:{&#34;eventToken&#34;:&#34;8npzxVENtQ1VNij1pCFlxp3BNhkKdJh9FdggP2851ydw54XAlJ3xNRImNE7jBWgsgJd3DNmZFmAmR22YP7jXhw==&#34;},&#34;correlationID&#34;:&#34;271513:0:0&#34;}}],&#34;propositionEventType&#34;:{&#34;display&#34;:1}}},&#34;eventType&#34;:&#34;decisioning.propositionDisplay&#34;,&#34;web&#34;:{&#34;webPageDetails&#34;:{&#34;URL&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;},&#34;webReferrer&#34;:{&#34;URL&#34;:&#34;&#34;}},&#34;device&#34;:{&#34;screenHeight&#34;:720,&#34;screenWidth&#34;:1280,&#34;screenOrientation&#34;:&#34;landscape&#34;},&#34;environment&#34;:{&#34;type&#34;:&#34;browser&#34;,&#34;browserDetails&#34;:{&#34;viewportWidth&#34;:1280,&#34;viewportHeight&#34;:720}},&#34;placeContext&#34;:{&#34;localTimezoneOffset&#34;:420,&#34;localTime&#34;:&#34;2025-05-08T11:31:19.691-07:00&#34;},&#34;timestamp&#34;:&#34;2025-05-08T18:31:19.691Z&#34;,&#34;implementationDetails&#34;:{&#34;name&#34;:&#34;https://ns.adobe.com/experience/alloy/reactor&#34;,&#34;version&#34;:&#34;2.26.0&#43;2.29.0&#34;,&#34;environment&#34;:&#34;browser&#34;}}}],&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true,&#34;entries&#34;:[{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_identity&#34;,&#34;value&#34;:&#34;CiYyODE4MDkyODQ4NjE4NzUyODY1MDM5MjUzNTg4NzAxMjI5MTU4OFISCJ2jqonrMhABGAEqA09SMjAA8AGdo6qJ6zI=&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_cluster&#34;,&#34;value&#34;:&#34;or2&#34;}]}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://c.contentsquare.net/pageview?ex=&amp;amp;dt=134&amp;amp;pvt=n&amp;amp;cvars=%7B%221%22%3A%5B%22Page%20Name%22%2C%22%2Fen-us%2F%22%5D%2C%222%22%3A%5B%22Site%20ID%22%2C%22www.crowdstrike.com%22%5D%2C%224%22%3A%5B%22URL%22%2C%22www.crowdstrike.com%2Fen-us%2F%22%5D%2C%2211%22%3A%5B%22DB%20Industry%20Data%22%2C%22%25demandbaseDataElement1%25%22%5D%2C%2212%22%3A%5B%22DB%20Company%20Data%22%2C%22%25demandbaseDataElement2%25%22%5D%7D&amp;amp;cvarp=%7B%221%22%3A%5B%22Page%20Name%22%2C%22%2Fen-us%2F%22%5D%2C%222%22%3A%5B%22Site%20ID%22%2C%22www.crowdstrike.com%22%5D%2C%224%22%3A%5B%22URL%22%2C%22www.crowdstrike.com%2Fen-us%2F%22%5D%2C%2211%22%3A%5B%22DB%20Industry%20Data%22%2C%22%25demandbaseDataElement1%25%22%5D%2C%2212%22%3A%5B%22DB%20Company%20Data%22%2C%22%25demandbaseDataElement2%25%22%5D%7D&amp;amp;la=en-US&amp;amp;uc=0&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;dr=&amp;amp;dw=1280&amp;amp;dh=7962&amp;amp;ww=1280&amp;amp;wh=720&amp;amp;sw=1280&amp;amp;sh=720&amp;amp;uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;amp;sn=1&amp;amp;hd=1746729079&amp;amp;v=15.85.1&amp;amp;pid=29632&amp;amp;pn=1&amp;amp;r=089855", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.google.com/pagead/1p-user-list/797629828/?random=1746729079496&amp;amp;cv=11&amp;amp;fst=1746727200000&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;data=event%3Dgtag.config&amp;amp;rfmt=3&amp;amp;fmt=3&amp;amp;is_vtc=1&amp;amp;cid=CAQSGwDZpuyz7viBNmNM2I-21Pt81hd1C8rcQz-rtg&amp;amp;random=246964701&amp;amp;rmt_tld=0&amp;amp;ipr=y", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.google.com/pagead/1p-user-list/952416460/?random=1746729079503&amp;amp;cv=11&amp;amp;fst=1746727200000&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;data=event%3Dgtag.config&amp;amp;rfmt=3&amp;amp;fmt=3&amp;amp;is_vtc=1&amp;amp;cid=CAQSGwDZpuyzPChu2ZLBUSUbxC76xLhpZO1EVFB39g&amp;amp;random=3015211708&amp;amp;rmt_tld=0&amp;amp;ipr=y", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://c.contentsquare.net/dvar?v=15.85.1&amp;amp;pid=29632&amp;amp;pn=1&amp;amp;sn=1&amp;amp;uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;amp;dv=H4sIAAAAAAAAA6tWSi72TSxJzsjMS%2FdOrVSyUjLQszQyMDUxNTUyMTE2NDIxNI83NDcxMzeyNDC3NDQ1VaoFAC3Wr880AAAA&amp;amp;ct=2&amp;amp;r=397738", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.linkedin.com/px/li_sync?redirect=https%3A%2F%2Fpx.ads.linkedin.com%2Fcollect%3Fv%3D2%26fmt%3Djs%26pid%3D64444%26time%3D1746729079483%26li_adsId%3Df8924343-8fd8-45cc-9a32-420fa0932ae3%26url%3Dhttps%253A%252F%252Fwww.crowdstrike.com%252Fen-us%252F%26cookiesTest%3Dtrue%26liSync%3Dtrue", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=4a16f21c-818e-4949-aaa6-19c62610ab56&amp;amp;batch_time=1746729079765", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079208,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;long_task&#34;:{&#34;id&#34;:&#34;106078a3-6ff5-4001-920f-36058e1a93fc&#34;,&#34;entry_type&#34;:&#34;long-task&#34;,&#34;duration&#34;:262000000},&#34;type&#34;:&#34;long_task&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079548,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;long_task&#34;:{&#34;id&#34;:&#34;49f17f73-87a0-4cef-8f55-917bca389815&#34;,&#34;entry_type&#34;:&#34;long-task&#34;,&#34;duration&#34;:61000000},&#34;type&#34;:&#34;long_task&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078135,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;ec1c968d-5957-4177-bbf1-07cefa43151e&#34;,&#34;type&#34;:&#34;document&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:159800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:225807,&#34;encoded_body_size&#34;:24785,&#34;decoded_body_size&#34;:225807,&#34;transfer_size&#34;:25085,&#34;download&#34;:{&#34;duration&#34;:7600000,&#34;start&#34;:152200000},&#34;first_byte&#34;:{&#34;duration&#34;:93400000,&#34;start&#34;:58800000},&#34;connect&#34;:{&#34;duration&#34;:56500000,&#34;start&#34;:2200000},&#34;ssl&#34;:{&#34;duration&#34;:37800000,&#34;start&#34;:20900000},&#34;dns&#34;:{&#34;duration&#34;:1400000,&#34;start&#34;:800000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078291,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;a17c1b8d-37eb-4cae-96ba-dabd0d9c83c8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.cookielaw.org/consent/bee15b7c-b632-450e-9003-9c8b60b3b978/OtAutoBlock.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:115100000,&#34;render_blocking_status&#34;:&#34;blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078291,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6a32a487-678c-43f3-84e4-2426e8cb1b76&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-analytics.lc-f643ec19011038440608ac5e1d5dd675-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:60400000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:37496,&#34;encoded_body_size&#34;:9470,&#34;decoded_body_size&#34;:37496,&#34;transfer_size&#34;:9770,&#34;download&#34;:{&#34;duration&#34;:2600000,&#34;start&#34;:57800000},&#34;first_byte&#34;:{&#34;duration&#34;:55800000,&#34;start&#34;:2000000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078291,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;33bb065e-0c15-4364-9501-35cd5b9b7eba&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.cookielaw.org/scripttemplates/otSDKStub.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:103800000,&#34;render_blocking_status&#34;:&#34;blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078291,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;7c948e1d-7c13-483e-b86f-0bc9caadc844&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/launch-d9bfd4283ab8.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:303900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:483821,&#34;encoded_body_size&#34;:131725,&#34;decoded_body_size&#34;:483821,&#34;transfer_size&#34;:132025,&#34;download&#34;:{&#34;duration&#34;:70000000,&#34;start&#34;:233900000},&#34;first_byte&#34;:{&#34;duration&#34;:28600000,&#34;start&#34;:205300000},&#34;connect&#34;:{&#34;duration&#34;:45400000,&#34;start&#34;:159600000},&#34;ssl&#34;:{&#34;duration&#34;:34300000,&#34;start&#34;:170700000},&#34;dns&#34;:{&#34;duration&#34;:23100000,&#34;start&#34;:136500000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078291,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;266cde8a-348b-4f09-b417-d465732de61c&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-language.lc-4418ee9ad99afc3c80f14f47d372a699-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:54500000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:609,&#34;encoded_body_size&#34;:363,&#34;decoded_body_size&#34;:609,&#34;transfer_size&#34;:663,&#34;download&#34;:{&#34;duration&#34;:400000,&#34;start&#34;:54100000},&#34;first_byte&#34;:{&#34;duration&#34;:51400000,&#34;start&#34;:2700000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078291,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e8139a81-87dd-4216-a620-b7296acc4608&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-base.lc-61e1f474009a3adcff6cc95dde2309a3-lc.min.css&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:56000000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:5108,&#34;encoded_body_size&#34;:2079,&#34;decoded_body_size&#34;:5108,&#34;transfer_size&#34;:2379,&#34;download&#34;:{&#34;duration&#34;:1300000,&#34;start&#34;:54700000},&#34;first_byte&#34;:{&#34;duration&#34;:52100000,&#34;start&#34;:2600000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078291,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1ef36eb5-5e88-4dc2-b4e8-11929d8d9847&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts.lc-475459b8ddb225de7aad82eaf0d465c8-lc.min.css&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:68900000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:12367,&#34;encoded_body_size&#34;:894,&#34;decoded_body_size&#34;:12367,&#34;transfer_size&#34;:1194,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:68700000},&#34;first_byte&#34;:{&#34;duration&#34;:66000000,&#34;start&#34;:2700000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078291,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;67914c64-203a-4fe4-8432-dd8a103f36dd&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-grid.lc-d75b4ada966a12a0acc4f4483174f716-lc.min.css&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:69000000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:103405,&#34;encoded_body_size&#34;:5386,&#34;decoded_body_size&#34;:103405,&#34;transfer_size&#34;:5686,&#34;download&#34;:{&#34;duration&#34;:2100000,&#34;start&#34;:66900000},&#34;first_byte&#34;:{&#34;duration&#34;:64600000,&#34;start&#34;:2300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;872c4097-8f8c-45ee-9188-2484fb08c068&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/clientlibs/granite/jquery.lc-f9e8e8c279baf6a1a278042afe4f395a-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:62400000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:101682,&#34;encoded_body_size&#34;:36477,&#34;decoded_body_size&#34;:101682,&#34;transfer_size&#34;:36777,&#34;download&#34;:{&#34;duration&#34;:2100000,&#34;start&#34;:60300000},&#34;first_byte&#34;:{&#34;duration&#34;:57600000,&#34;start&#34;:2700000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;85b5e9a9-1eba-4f67-be0d-91545cf2047b&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-dependencies.lc-2929dd9d69a653da2cf8fe429017e6b6-lc.min.css&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:64600000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:13612,&#34;encoded_body_size&#34;:2634,&#34;decoded_body_size&#34;:13612,&#34;transfer_size&#34;:2934,&#34;download&#34;:{&#34;duration&#34;:300000,&#34;start&#34;:64300000},&#34;first_byte&#34;:{&#34;duration&#34;:61500000,&#34;start&#34;:2800000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e892957c-8732-470d-818b-ccf83cc0e6ec&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-dependencies.lc-bb1dbdd53e32240429436d11ecb5f036-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:67000000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:136434,&#34;encoded_body_size&#34;:40484,&#34;decoded_body_size&#34;:136434,&#34;transfer_size&#34;:40784,&#34;download&#34;:{&#34;duration&#34;:600000,&#34;start&#34;:66400000},&#34;first_byte&#34;:{&#34;duration&#34;:63600000,&#34;start&#34;:2800000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;dde666bf-1388-4d0a-bcd6-00993bb2643b&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-common.lc-d2017b1bb6b94dc55ac836b3e60fc92d-lc.min.css&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:80900000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:385333,&#34;encoded_body_size&#34;:35709,&#34;decoded_body_size&#34;:385333,&#34;transfer_size&#34;:36009,&#34;download&#34;:{&#34;duration&#34;:10500000,&#34;start&#34;:70400000},&#34;first_byte&#34;:{&#34;duration&#34;:67500000,&#34;start&#34;:2900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0,&#34;start_session_replay_recording_manually&#34;:true},&#34;document_version&#34;:1,&#34;page_states&#34;:[{&#34;state&#34;:&#34;active&#34;,&#34;start&#34;:1615600000}]},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078135,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;,&#34;sampled_for_replay&#34;:false},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;,&#34;action&#34;:{&#34;count&#34;:0},&#34;frustration&#34;:{&#34;count&#34;:0},&#34;cumulative_layout_shift&#34;:0,&#34;error&#34;:{&#34;count&#34;:0},&#34;is_active&#34;:true,&#34;loading_type&#34;:&#34;initial_load&#34;,&#34;long_task&#34;:{&#34;count&#34;:0},&#34;resource&#34;:{&#34;count&#34;:0},&#34;time_spent&#34;:1616000000},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;type&#34;:&#34;view&#34;,&#34;privacy&#34;:{&#34;replay_level&#34;:&#34;mask-user-input&#34;}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=49e0090f-932c-4c65-92ec-2abfc4ee1ddf&amp;amp;batch_time=1746729079767", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;7cd1511e-4106-444c-8589-e3cbb01cbe7b&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-header.lc-0b9b1c99ad703dffc489ae2fdecad2b2-lc.min.css&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:118900000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:126447,&#34;encoded_body_size&#34;:14263,&#34;decoded_body_size&#34;:126447,&#34;transfer_size&#34;:14563,&#34;download&#34;:{&#34;duration&#34;:2100000,&#34;start&#34;:116800000},&#34;first_byte&#34;:{&#34;duration&#34;:113900000,&#34;start&#34;:2900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b44bca58-7a01-4563-aa88-c18382e98754&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/clientlibs/granite/utils.lc-899004cc02c33efc1f6694b1aee587fd-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:121200000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:8529,&#34;encoded_body_size&#34;:3521,&#34;decoded_body_size&#34;:8529,&#34;transfer_size&#34;:3821,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:121000000},&#34;first_byte&#34;:{&#34;duration&#34;:118200000,&#34;start&#34;:2800000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;54ce2b60-3b4e-4451-9ef5-e3771358a575&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-addsearch.lc-fa338b27fbcf1b58755f6bc3391495d1-lc.min.css&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:64100000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:8620,&#34;encoded_body_size&#34;:2650,&#34;decoded_body_size&#34;:8620,&#34;transfer_size&#34;:2950,&#34;download&#34;:{&#34;duration&#34;:1700000,&#34;start&#34;:62400000},&#34;first_byte&#34;:{&#34;duration&#34;:59500000,&#34;start&#34;:2900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b6d610e1-8b73-47be-8ce6-8b53fd1d91f9&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-dotcom.lc-33079283bcb9b23ffc4ce228005513c1-lc.min.css&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:109500000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:653280,&#34;encoded_body_size&#34;:64992,&#34;decoded_body_size&#34;:653280,&#34;transfer_size&#34;:65292,&#34;download&#34;:{&#34;duration&#34;:2800000,&#34;start&#34;:106700000},&#34;first_byte&#34;:{&#34;duration&#34;:103800000,&#34;start&#34;:2900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;33fb2b08-061c-4312-93d9-ff9a8a05db6f&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/graphics-and-illustrations/marketplace/addsearch-icon.svg&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:225200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:528,&#34;encoded_body_size&#34;:255,&#34;decoded_body_size&#34;:528,&#34;transfer_size&#34;:555,&#34;download&#34;:{&#34;duration&#34;:11200000,&#34;start&#34;:214000000},&#34;first_byte&#34;:{&#34;duration&#34;:76700000,&#34;start&#34;:137300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6ba907e2-e8bd-4b9d-8a4f-ded8935076a7&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/shopping-cart-empty&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:151400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078292,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;3575ccd5-a578-44bd-ad37-f9463fc00b80&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/login-icon&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:157100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078293,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;740b7594-9512-4cad-a8f9-5cbd82c2f6f9&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/.rum/@adobe/helix-rum-js@%5E2/src/index.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:79500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:4648,&#34;encoded_body_size&#34;:1996,&#34;decoded_body_size&#34;:4648,&#34;transfer_size&#34;:2296,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:79300000},&#34;first_byte&#34;:{&#34;duration&#34;:77600000,&#34;start&#34;:1700000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078294,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9ea1c67d-981a-49ad-b153-96faa0793a4d&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/experience-breach-Icon&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:162600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078294,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9dfc5d5d-8e49-4dd4-87ff-ed35fda32baf&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/empty-cart-image&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:186400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078294,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;dfce9dea-f66a-467b-a636-c9ac324abbee&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/phone-icon&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:189200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078294,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;cd68964a-3a45-4399-ac4d-41c0306cc167&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://play.vidyard.com/72xzgRJUfHEGrXeQaYMY5M.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:494000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078294,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;37def5de-b95f-4b61-ab35-270facd99f9d&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/adversaries/open-report.png&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:326900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:2240974,&#34;encoded_body_size&#34;:2240974,&#34;decoded_body_size&#34;:2240974,&#34;transfer_size&#34;:2241274,&#34;download&#34;:{&#34;duration&#34;:115700000,&#34;start&#34;:211200000},&#34;first_byte&#34;:{&#34;duration&#34;:76500000,&#34;start&#34;:134700000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078294,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c321fb21-6cce-4458-9a4b-6da8bd9ee041&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://play.vidyard.com/PBjCLnZ26eK9VG3ifwnnVp.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:661500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;41040fe1-eca3-4bf6-9eae-f3bb7d60cdb0&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/adversaries/adversary-group.png&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:291200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:777535,&#34;encoded_body_size&#34;:777535,&#34;decoded_body_size&#34;:777535,&#34;transfer_size&#34;:777835,&#34;download&#34;:{&#34;duration&#34;:63600000,&#34;start&#34;:227600000},&#34;first_byte&#34;:{&#34;duration&#34;:93000000,&#34;start&#34;:134600000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6da356ba-f940-496b-88f4-823ce34505ca&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/adversaries/hi-res-adv-bg.png&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:337600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:2081896,&#34;encoded_body_size&#34;:2081896,&#34;decoded_body_size&#34;:2081896,&#34;transfer_size&#34;:2082196,&#34;download&#34;:{&#34;duration&#34;:126000000,&#34;start&#34;:211600000},&#34;first_byte&#34;:{&#34;duration&#34;:77000000,&#34;start&#34;:134600000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;04c3bfa9-1ed8-4f19-a193-c953baa13147&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://play.vidyard.com/SJBQ6ji2tUz4syr4dz2y3s.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:774900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=b84d299e-9c6e-4cff-a8ee-bc89ac19666b&amp;amp;batch_time=1746729079767", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;f9e8f0c0-9b18-4e5c-abf6-6ed5f7f8a09b&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://play.vidyard.com/nTrNLYDyFzyH9dBgveMv91.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:795700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;3fbe8234-a64b-46c8-8f53-ad7cdfbaa2b2&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/privacyoptions-icon-footer&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:162900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;25fc196c-6c61-42d6-bdf8-08a22b2b8dbd&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-addsearch.lc-b8f6141a81735ad2112829f5c5025005-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:219500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:245273,&#34;encoded_body_size&#34;:67809,&#34;decoded_body_size&#34;:245273,&#34;transfer_size&#34;:68109,&#34;download&#34;:{&#34;duration&#34;:8600000,&#34;start&#34;:210900000},&#34;first_byte&#34;:{&#34;duration&#34;:77300000,&#34;start&#34;:133600000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;53289a9d-ef7b-478a-a3a4-f9cb28f390b4&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-header.lc-13360fa8c738773238c6e841d8970e35-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:221700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:251421,&#34;encoded_body_size&#34;:68027,&#34;decoded_body_size&#34;:251421,&#34;transfer_size&#34;:68327,&#34;download&#34;:{&#34;duration&#34;:11100000,&#34;start&#34;:210600000},&#34;first_byte&#34;:{&#34;duration&#34;:77200000,&#34;start&#34;:133400000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8bcbc92d-7e27-4808-b457-f3f71ecf08b6&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-common.lc-9bb6aad5b81df0110105b0a332fb1ecd-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:235100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1107126,&#34;encoded_body_size&#34;:616414,&#34;decoded_body_size&#34;:1107126,&#34;transfer_size&#34;:616714,&#34;download&#34;:{&#34;duration&#34;:46400000,&#34;start&#34;:188700000},&#34;first_byte&#34;:{&#34;duration&#34;:55800000,&#34;start&#34;:132900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;20ab8b89-c40b-4b71-9129-0edbab202429&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-animation.lc-012003c5a08ba65bfd2d6d6b777b7f4e-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:210200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:161437,&#34;encoded_body_size&#34;:61321,&#34;decoded_body_size&#34;:161437,&#34;transfer_size&#34;:61621,&#34;download&#34;:{&#34;duration&#34;:10400000,&#34;start&#34;:199800000},&#34;first_byte&#34;:{&#34;duration&#34;:65800000,&#34;start&#34;:134000000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;07e58a0e-d327-4ebd-bb42-c44186dad5ef&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/core/wcm/components/commons/site/clientlibs/container.lc-0a6aff292f5cc42142779cde92054524-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:218400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1271,&#34;encoded_body_size&#34;:491,&#34;decoded_body_size&#34;:1271,&#34;transfer_size&#34;:791,&#34;download&#34;:{&#34;duration&#34;:7400000,&#34;start&#34;:211000000},&#34;first_byte&#34;:{&#34;duration&#34;:77100000,&#34;start&#34;:133900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6480e82c-db2b-4ef1-8947-4c379ef32aa8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-base.lc-98b44ec74775c5bc76b0744df1c9b66c-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:218500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:75430,&#34;encoded_body_size&#34;:22502,&#34;decoded_body_size&#34;:75430,&#34;transfer_size&#34;:22802,&#34;download&#34;:{&#34;duration&#34;:7700000,&#34;start&#34;:210800000},&#34;first_byte&#34;:{&#34;duration&#34;:76900000,&#34;start&#34;:133900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e420a3d7-d696-49b6-89f1-a8f0d7b62a74&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-lottie.lc-59169d2e9977544fd3ae80d885311bce-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:221800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:365175,&#34;encoded_body_size&#34;:93007,&#34;decoded_body_size&#34;:365175,&#34;transfer_size&#34;:93307,&#34;download&#34;:{&#34;duration&#34;:21800000,&#34;start&#34;:200000000},&#34;first_byte&#34;:{&#34;duration&#34;:66000000,&#34;start&#34;:134000000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078295,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e0b699d5-92b8-4504-a049-7ca6f489fd7b&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-dotcom.lc-db18d2853624d228ab184335c97a1302-lc.min.js&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:263600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1050437,&#34;encoded_body_size&#34;:354684,&#34;decoded_body_size&#34;:1050437,&#34;transfer_size&#34;:354984,&#34;download&#34;:{&#34;duration&#34;:52800000,&#34;start&#34;:210800000},&#34;first_byte&#34;:{&#34;duration&#34;:77000000,&#34;start&#34;:133800000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078424,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;363ba4e4-e284-4f85-9f5c-0191696979a1&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://use.typekit.net/zya3koo.css&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:95000000,&#34;render_blocking_status&#34;:&#34;blocking&#34;,&#34;size&#34;:2715,&#34;encoded_body_size&#34;:692,&#34;decoded_body_size&#34;:2715,&#34;transfer_size&#34;:992,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:94900000},&#34;first_byte&#34;:{&#34;duration&#34;:22000000,&#34;start&#34;:72900000},&#34;connect&#34;:{&#34;duration&#34;:52000000,&#34;start&#34;:20200000},&#34;ssl&#34;:{&#34;duration&#34;:37800000,&#34;start&#34;:34400000},&#34;dns&#34;:{&#34;duration&#34;:20000000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078612,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d7104df3-c2a4-4893-a9bd-b8d657fbd737&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/backgrounds/falcon-platform-texture-bg-v2.png&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:103800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:711265,&#34;encoded_body_size&#34;:711265,&#34;decoded_body_size&#34;:711265,&#34;transfer_size&#34;:711565,&#34;download&#34;:{&#34;duration&#34;:10600000,&#34;start&#34;:93200000},&#34;first_byte&#34;:{&#34;duration&#34;:93000000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078614,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;01946aaf-8a0a-491f-ab1a-b45a215fc514&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/sans/Crowdstrike_Sharp_Sans/CrowdstrikeSharpSans-Mdm.woff2&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:81100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:0,&#34;encoded_body_size&#34;:0,&#34;decoded_body_size&#34;:0,&#34;transfer_size&#34;:300,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:80900000},&#34;first_byte&#34;:{&#34;duration&#34;:67900000,&#34;start&#34;:13000000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078614,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;388e5137-4675-4206-b2b4-d2b1afecf102&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/neue_haas_grot_disp_pro/NeueHaasDisplayLight.ttf&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:84300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:101684,&#34;encoded_body_size&#34;:37049,&#34;decoded_body_size&#34;:101684,&#34;transfer_size&#34;:37349,&#34;download&#34;:{&#34;duration&#34;:400000,&#34;start&#34;:83900000},&#34;first_byte&#34;:{&#34;duration&#34;:70900000,&#34;start&#34;:13000000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078614,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;dad8aeff-d0a0-4744-945f-7be122ef6b38&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/haas_grot_disp/HaasGrotDisp-65Medium.woff2&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:102100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:32876,&#34;encoded_body_size&#34;:32904,&#34;decoded_body_size&#34;:32876,&#34;transfer_size&#34;:33204,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:101900000},&#34;first_byte&#34;:{&#34;duration&#34;:89100000,&#34;start&#34;:12800000}},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=01622766-3828-45b8-98a6-86bd52dcbaae&amp;amp;batch_time=1746729079768", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078614,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;7ef1d040-62cf-435c-8548-6b5d1e857eff&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/neue_haas_grot_disp_pro/NeueHaasDisplayMedium.ttf&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:89300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:101948,&#34;encoded_body_size&#34;:36626,&#34;decoded_body_size&#34;:101948,&#34;transfer_size&#34;:36926,&#34;download&#34;:{&#34;duration&#34;:1100000,&#34;start&#34;:88200000},&#34;first_byte&#34;:{&#34;duration&#34;:75300000,&#34;start&#34;:12900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078615,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6d7692f7-8f6a-42f8-858f-6e1c5b8d7c49&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/sans/Crowdstrike_Sharp_Sans/CrowdstrikeSharpSans-Bold.woff2&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:83700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:0,&#34;encoded_body_size&#34;:0,&#34;decoded_body_size&#34;:0,&#34;transfer_size&#34;:300,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:83600000},&#34;first_byte&#34;:{&#34;duration&#34;:70800000,&#34;start&#34;:12800000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078615,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8033edd0-1677-4f8f-ab93-3938b034efb4&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/haas_grot_disp/HaasGrotDisp-55Roman.woff2&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:101600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:32272,&#34;encoded_body_size&#34;:32295,&#34;decoded_body_size&#34;:32272,&#34;transfer_size&#34;:32595,&#34;download&#34;:{&#34;duration&#34;:300000,&#34;start&#34;:101300000},&#34;first_byte&#34;:{&#34;duration&#34;:88400000,&#34;start&#34;:12900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078615,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;668da750-ad6d-4820-8cdd-a0da8eb3574b&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/haas_grot_disp/HaasGrotDisp-45Light.woff2&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:33280,&#34;encoded_body_size&#34;:33308,&#34;decoded_body_size&#34;:33280,&#34;transfer_size&#34;:33608,&#34;download&#34;:{&#34;duration&#34;:1300000,&#34;start&#34;:81400000},&#34;first_byte&#34;:{&#34;duration&#34;:68500000,&#34;start&#34;:12900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078626,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5b8eda25-b627-43bb-9aae-b57acd14d2a4&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/image/crowdstrikeinc/25-SRV-007_Forrester%20Wave%20MDR%203153x3860?ts=1746621981959&amp;dpr=off&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:40800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078626,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;f9cc08b8-ab49-45d6-9f24-f55976d5f274&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Platform-4-16-25-darkmode-HP?ts=1746621981798&amp;dpr=off&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:66600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078626,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;255330cc-0da4-40e6-bf80-24f714c5d7a9&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/black-primary-crowdstrike-logo-1-addedPadding-3?ts=1746646455211&amp;dpr=off&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:254400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078626,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;09477828-7c24-4bec-9c1f-1c4f0068cf9e&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/image/crowdstrikeinc/identity-security-services-new?ts=1746621982825&amp;dpr=off&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:86400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078626,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;bcccb87d-5993-4a0e-b01a-971daa2e3b8d&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/image/crowdstrikeinc/cloud-lock-generic-img?ts=1746621982750&amp;dpr=off&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078626,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;de610d5a-00dc-4548-a9ef-0267cf089d34&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/image/crowdstrikeinc/soc-survival-guide-2.bak?ts=1746621982899&amp;dpr=off&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:95000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078678,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;66ed4d8b-d2cf-4922-bb5c-90b51372186b&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.cookielaw.org/scripttemplates/202401.2.0/otBannerSdk.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:41300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078692,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;2e5db998-2ddf-4a4e-9d6e-ba348064d505&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/image/crowdstrikeinc/try-free-footer&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:38300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078697,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;3a263451-197f-48c0-983a-a8f3807fcbd5&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/image/crowdstrikeinc/charlotte-hp-hero-1?&amp;hei=769&amp;wid=1600&amp;qlt=90&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:35800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078713,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;40eef29f-e9e7-40f2-8290-292cd4e32ceb&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/sans/Crowdstrike_Sharp_Sans/CrowdstrikeSharpSans-Mdm.ttf&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:63500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:129260,&#34;encoded_body_size&#34;:59406,&#34;decoded_body_size&#34;:129260,&#34;transfer_size&#34;:59706,&#34;download&#34;:{&#34;duration&#34;:700000,&#34;start&#34;:62800000},&#34;first_byte&#34;:{&#34;duration&#34;:62000000,&#34;start&#34;:800000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078713,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c260eedc-668b-45d7-834c-af9d87f9a7e0&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/sans/Crowdstrike_Sharp_Sans/CrowdstrikeSharpSans-Bold.ttf&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:64800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:132096,&#34;encoded_body_size&#34;:60007,&#34;decoded_body_size&#34;:132096,&#34;transfer_size&#34;:60307,&#34;download&#34;:{&#34;duration&#34;:1700000,&#34;start&#34;:63100000},&#34;first_byte&#34;:{&#34;duration&#34;:62400000,&#34;start&#34;:700000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078722,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8cec6ce6-78e5-4b9a-a804-6fb93deb5959&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/image/crowdstrikeinc/gartner-mq-leader-report?ts=1746621982064&amp;dpr=off&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:30000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078722,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;20342cd0-54fc-42c1-8638-c2af994a87e2&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/image/crowdstrikeinc/frost-radar-2024-2?ts=1746621982168&amp;dpr=off&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:39700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=bade6cb7-6b23-48ad-a911-0146c769366b&amp;amp;batch_time=1746729079768", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6dd32be7-aea3-4cce-a9c9-5082a9c9445c&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Endpoint-Security-2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:80500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c375f7c7-bbea-4ed4-ac14-5129673bcdbe&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Exposure-Management-2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:72600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;dd3a77a8-2f6f-4adb-8c30-44e530fca2e4&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/SaaS-Security&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:73700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8852ef0f-304b-4b05-9da4-90c7c8df9f24&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/IT-Automation&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:73900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;79b74129-cac9-4875-9fb0-ac130ea5b1da&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Threat-Intelligence-Hunting&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;7bbfea89-4398-4821-ae16-f744c39d6888&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Cloud-Security-2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;47a70fad-e02a-4507-9f4b-25afcc875db7&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Identity-Protection-1&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:80600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6e8880c8-a142-4d31-978f-ff6d227d263f&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Workflow-Automation&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:81700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9c810f39-431f-4b6f-8f0e-a99c9a44895d&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Generative-AI&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:179400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1c0b3634-2c20-4dfc-8c4b-1592c0a21c21&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Next-Gen-SIEM&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;0f21ef38-8197-4675-95a7-972f05b7e537&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Data-Protection-3&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d96f2ffa-4ff1-48cd-8d9d-17f858ee00f6&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/xiot-security&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078885,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;13af036e-3259-4624-a0b4-e4d23f9303fd&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/launch-d9bfd4283ab8.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;cache&#34;,&#34;duration&#34;:900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:483821,&#34;encoded_body_size&#34;:131725,&#34;decoded_body_size&#34;:483821,&#34;transfer_size&#34;:0,&#34;download&#34;:{&#34;duration&#34;:700000,&#34;start&#34;:200000},&#34;first_byte&#34;:{&#34;duration&#34;:0,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078887,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;eb7c872a-3c4a-4845-b9e3-fb376b3ff31b&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.cookielaw.org/logos/static/ot_close.svg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:36800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078896,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5d41ca59-05a3-4e48-bbda-ebfd9f324bb0&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.cookielaw.org/logos/static/powered_by_logo.svg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:35800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078896,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;51a17285-52df-4ccb-9e95-0004b167edb1&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.cookielaw.org/logos/c109dae9-46f3-4e91-a59e-7844ef645107/a4486e7b-87d4-4354-a844-b53722d937ac/e2865569-21ec-4213-8a1d-4f0ed5a1672e/CS_Logo_2022_In-Line_All-Red_RGB_(1).png&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:37600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078904,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9c6ac926-cb43-4cf4-aa9d-39e580089bbd&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://connect.facebook.net/en_US/fbevents.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:215300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:278083,&#34;encoded_body_size&#34;:70063,&#34;decoded_body_size&#34;:278083,&#34;transfer_size&#34;:70363,&#34;download&#34;:{&#34;duration&#34;:95400000,&#34;start&#34;:119900000},&#34;first_byte&#34;:{&#34;duration&#34;:47500000,&#34;start&#34;:72400000},&#34;connect&#34;:{&#34;duration&#34;:69300000,&#34;start&#34;:1800000},&#34;ssl&#34;:{&#34;duration&#34;:41100000,&#34;start&#34;:30000000},&#34;dns&#34;:{&#34;duration&#34;:1600000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078916,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;71c7a90b-3100-47d9-9285-e6d1e2923466&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://www.googletagmanager.com/gtag/js?id=G-ZKTET1D58V&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:232500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078918,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d9c9f86b-52b1-4204-8b5b-ac95289bf0e8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://js.driftt.com/include/1746729300000/9d4udx6ceimp.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:330400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	
	req.Header.Set("Accept-Language", "en-US")
	
	req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != -1 {
		t.Errorf("Expected status %d, got %d", -1, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Axhr%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=4aa41092-e0c8-4370-836d-65f68d90aee3&amp;amp;batch_time=1746729079770", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6dd32be7-aea3-4cce-a9c9-5082a9c9445c&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Endpoint-Security-2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:80500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c375f7c7-bbea-4ed4-ac14-5129673bcdbe&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Exposure-Management-2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:72600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;dd3a77a8-2f6f-4adb-8c30-44e530fca2e4&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/SaaS-Security&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:73700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8852ef0f-304b-4b05-9da4-90c7c8df9f24&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/IT-Automation&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:73900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;79b74129-cac9-4875-9fb0-ac130ea5b1da&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Threat-Intelligence-Hunting&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;7bbfea89-4398-4821-ae16-f744c39d6888&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Cloud-Security-2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;47a70fad-e02a-4507-9f4b-25afcc875db7&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Identity-Protection-1&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:80600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6e8880c8-a142-4d31-978f-ff6d227d263f&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Workflow-Automation&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:81700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9c810f39-431f-4b6f-8f0e-a99c9a44895d&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Generative-AI&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:179400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1c0b3634-2c20-4dfc-8c4b-1592c0a21c21&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Next-Gen-SIEM&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;0f21ef38-8197-4675-95a7-972f05b7e537&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/Data-Protection-3&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078753,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d96f2ffa-4ff1-48cd-8d9d-17f858ee00f6&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://assets.crowdstrike.com/is/content/crowdstrikeinc/xiot-security&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:74800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078885,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;13af036e-3259-4624-a0b4-e4d23f9303fd&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/launch-d9bfd4283ab8.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;cache&#34;,&#34;duration&#34;:900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:483821,&#34;encoded_body_size&#34;:131725,&#34;decoded_body_size&#34;:483821,&#34;transfer_size&#34;:0,&#34;download&#34;:{&#34;duration&#34;:700000,&#34;start&#34;:200000},&#34;first_byte&#34;:{&#34;duration&#34;:0,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078887,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;eb7c872a-3c4a-4845-b9e3-fb376b3ff31b&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.cookielaw.org/logos/static/ot_close.svg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:36800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078896,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5d41ca59-05a3-4e48-bbda-ebfd9f324bb0&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.cookielaw.org/logos/static/powered_by_logo.svg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:35800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078896,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;51a17285-52df-4ccb-9e95-0004b167edb1&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.cookielaw.org/logos/c109dae9-46f3-4e91-a59e-7844ef645107/a4486e7b-87d4-4354-a844-b53722d937ac/e2865569-21ec-4213-8a1d-4f0ed5a1672e/CS_Logo_2022_In-Line_All-Red_RGB_(1).png&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:37600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078904,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9c6ac926-cb43-4cf4-aa9d-39e580089bbd&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://connect.facebook.net/en_US/fbevents.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:215300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:278083,&#34;encoded_body_size&#34;:70063,&#34;decoded_body_size&#34;:278083,&#34;transfer_size&#34;:70363,&#34;download&#34;:{&#34;duration&#34;:95400000,&#34;start&#34;:119900000},&#34;first_byte&#34;:{&#34;duration&#34;:47500000,&#34;start&#34;:72400000},&#34;connect&#34;:{&#34;duration&#34;:69300000,&#34;start&#34;:1800000},&#34;ssl&#34;:{&#34;duration&#34;:41100000,&#34;start&#34;:30000000},&#34;dns&#34;:{&#34;duration&#34;:1600000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078916,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;71c7a90b-3100-47d9-9285-e6d1e2923466&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://www.googletagmanager.com/gtag/js?id=G-ZKTET1D58V&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:232500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078918,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d9c9f86b-52b1-4204-8b5b-ac95289bf0e8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://js.driftt.com/include/1746729300000/9d4udx6ceimp.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:330400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://rtp-static.marketo.com/rtp/libs/jqueryui/1.13.2/jquery-custom-ui.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://px.ads.linkedin.com/collect?v=2&amp;amp;fmt=js&amp;amp;pid=64444&amp;amp;time=1746729079483&amp;amp;li_adsId=f8924343-8fd8-45cc-9a32-420fa0932ae3&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;cookiesTest=true&amp;amp;liSync=true", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://web.chtbl.com/track", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;uid&#34;:&#34;ca2f6b2f-50c2-4d89-ab10-6af1d38fd2f3&#34;,&#34;pvid&#34;:&#34;c6dfd647-d519-462b-b7d6-d7b1ca3d0018&#34;,&#34;sid&#34;:&#34;64e43b05-41aa-49f6-83a3-9026f1f4896c&#34;,&#34;action&#34;:&#34;init&#34;,&#34;id&#34;:&#34;12175115ee0e8ce7aa41bcd5d4a663a5&#34;,&#34;context&#34;:{&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;},&#34;properties&#34;:{}}`))
	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != -1 {
		t.Errorf("Expected status %d, got %d", -1, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/msg?a=2&amp;amp;sid=crowdstrike-1746729079515-66d02f45&amp;amp;aid=crowdstrike&amp;amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;viewedTypes=&amp;amp;0.5769826221088862&amp;amp;rts=1746729079864", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://px4.ads.linkedin.com/collect?v=2&amp;amp;fmt=js&amp;amp;pid=64444&amp;amp;time=1746729079483&amp;amp;li_adsId=f8924343-8fd8-45cc-9a32-420fa0932ae3&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;cookiesTest=true&amp;amp;liSync=true&amp;amp;e_ipv6=AQKE_-sKYAL0lAAAAZaxKpRv1TXChQQeIgJ3DwVULni7wGGVjYVUBbweB149tx2kdc5zSAkjN_Fqyon8IFsr9Pw0aRjGrw", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/widgetapp/2025-05-08-12-13-34/widget_app_base_1746706414131.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://px.ads.linkedin.com/wa/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;pids&#34;:[64444],&#34;scriptVersion&#34;:213,&#34;time&#34;:1746729080058,&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;url&#34;:&#34;https://crowdstrike.com/en-us/&#34;,&#34;pageTitle&#34;:&#34;CrowdStrike: We Stop Breaches with AI-native Cybersecurity&#34;,&#34;websiteSignalRequestId&#34;:&#34;e727d704-7d05-8faa-325a-5b1ab056279f&#34;,&#34;isTranslated&#34;:false,&#34;liFatId&#34;:&#34;&#34;,&#34;liGiant&#34;:&#34;&#34;,&#34;misc&#34;:{&#34;psbState&#34;:-4},&#34;isLinkedInApp&#34;:false,&#34;hem&#34;:null,&#34;signalType&#34;:&#34;PAGE_VISIT&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://bat.bing.com/bat.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.googleadservices.com/pagead/conversion/797629828/?random=1746729080069&amp;amp;cv=11&amp;amp;fst=1746729080069&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gcl_ctr=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;label=hozuCPn52LoYEIS7q_wC&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;gtm_ee=1&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;capi=1&amp;amp;data=event%3Dconversion&amp;amp;rfmt=3&amp;amp;fmt=4", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://td.doubleclick.net/td/rul/797629828?random=1746729080069&amp;amp;cv=11&amp;amp;fst=1746729080069&amp;amp;fmt=3&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gcl_ctr=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;label=hozuCPn52LoYEIS7q_wC&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;gtm_ee=1&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;capi=1&amp;amp;data=event%3Dconversion&amp;amp;ct_cookie_present=0", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.googleadservices.com/pagead/conversion/797629828/?label=hozuCPn52LoYEIS7q_wC&amp;amp;guid=ON&amp;amp;script=0", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://obs.fishrobotflower.com/tracker/tc_imp.gif?e=37dfbd8ee84e00126ee8c037e347829d9225c24f567d43d6da1908be6245cad7bd70a976750ef80ed89373bfe70e9c20c1e53e8d5a138c6c2717071a10acf9f29f671a82d589562b6d1ef878215780338a33c103300376c5065359600d5ccfbf394f77be26bb25cb43e2913df05365ac0b297d1bdc57e446f492d7da3dbb280ef67ccaa8073f8d0e3717224793845731f360b3f493a0180dec1edae97dfa2bc8169b1adc597cff3200e714561c4b92177af998ffe4198b6dec06c213f85e162ae7d133722b325f817c99ec59b058609fc6e359143e3dd385293e88864c06513c157a77bb9e70392652b48d1c2ad7f4ec3ee3b8192d4079b4a7a6928677a0dadf5ae8489e5d2e019cbecbf7af2b95dfe57594351ccdeb8b795904fd7367a095166dbdd0aa73d902fd12b0801d91499d9978953b67919d27dc6796d3bbe193fdbd4c38fc28b0bdfd354371fe8f719aa61af7010642dd4245c2979684c0ec8825ce3bc68ebe817fa27d763bc97f27db838003679ea875ac41bfc003d667a87e5346588c6eee5f8b89ee7ef8df78f02da7c19cdb27d60a0a2dea5d540b28073ed7ab67b355c4cb9764fdde25b3d02dd0c2b1be7c85b7f409746c89a91004d7b41d5107f7fae939fa2890f802fdfc3ce52dd36dcc068a50aefc2ecd0823e2ac90633f3d6402340a5b13ace68e65d205be5752517faac71386cd464887eccdb0ae6ddcece3f880ffadf84b4835fe4586a85e6b844a601fd16a59746ef07aa6771ec1f9dd71c83cc52fe0a82cc8a40e99721322210760129cc0e4db207a559a184689d2d9bbb0c24f049eb0be7f8ad2757adbce0f3a68fd8dc45020b23b98599e1cff22981ad8d40efbb337fe868bd49e5b7feee3337f9f71a68aacb6119d6d280b02956a5f4db6a25b93f112ca9887d552cd839b2e083745d321b951ef0ce6936384335d20463e7b78e3538e863a4bcf1cc7c2f03da8cb0be60d76e678072c8df4c3491fb725e78402db1930074eba3c2e8bafd061d52c49e4ff15ebbd9d88e11aa90924605fe06b3716d9c9584eb02ee6e0c02a8bd5eb0c547ca5633e507e867459e6c8e66c40cd260bbbd6c353d5c6f88f0f30b05bd61b781488173175d4cc1615da92b2aabde142a563f930b77096d65d9afff22a2bfdd215808cb4b1c4f6cadf741fddf42f54872dd257c81cf3ee28a8cad1c731efc4c7f2fb5a69cba2b5af51cb5594404921413d88823335dc60f64f8ac965ef0522e2714c9766e69c7575b97521525796f3d12f87c74c603d1e6ac2d583ffce748a07e2ef4b7ee4de495b5018376e2de975e2a73802b3&amp;amp;cri=jfECv66HI9&amp;amp;ts=429&amp;amp;cb=1746729080075", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://api.userway.org/api/v1/tunings/dyvvHf6oG0", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;s&#34;:{},&#34;o&#34;:&#34;https://www.crowdstrike.com&#34;,&#34;uid&#34;:&#34;52e86a39-bbce-455f-867d-40c87b1494a2&#34;,&#34;v&#34;:&#34;2025-05-08-12-13-34&#34;,&#34;c&#34;:&#34;&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/u?mapType=mkto&amp;amp;mapValue=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729080105&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=1&amp;amp;a=crowdstrike.com&amp;amp;rnd=598845&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729080105", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://c.contentsquare.net/dvar?v=15.85.1&amp;amp;pid=29632&amp;amp;pn=1&amp;amp;sn=1&amp;amp;uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;amp;dv=H4sIAAAAAAAAA3WSX0%2BDMBTFv0pDfBwG5uZ0b4VtkWR%2FklUT4wupcMFGaElbtqHxu9tmOmDTR%2Fo79%2FScWz4dHMT4MX4QJVQ0B7SInudkkLGDu1favRl5rt%2F7GjpTpy%2Bezg8VSAY8AYT7k%2F%2BjYQ85AxNjFd%2BSeCZKyri547WWOoMiLSm%2FTkTZKkJRcy0bFJENCkUKRvtEWow5r2mBtrADXhs4nkyGnued%2BBpHITEz45Hv%2B57TP0czUIlklWbCZthkGUtAIZGhJd03IFWrJ1FoFHfGpHd25rCE3IQhIHfWqFuirChv0JqWtkBgyqLFse1FU1uQMw0pIprqngvTlpJEaK1SWkAnipUahiX7EJyegd%2FF4ZcWzMuqEA1YZq51pv4l2VKe2zEPuei%2BnYx4Wqtj0KBWjINSf1T%2BeZGTydV4ZWyufG%2FV1eTHva2F1G8gOcKl%2BUmSTv5IoaCgyXvBlFmJkWa0UOB8fQN6Gti9yQIAAA%3D%3D&amp;amp;ct=2&amp;amp;r=414786", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=fd5f0583-5464-432e-b56e-743069cb66d4&amp;amp;batch_time=1746729080201", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078919,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;a24cf75a-e1f8-4a89-b3b0-89f244b687cd&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/scripts/bizible.js?account=crowdstrike.com&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:182900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078919,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8438408b-dc0a-4c5c-8295-960f9e24bda2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCaf408a7d7bb64bba9c9ebd33e91577a5-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:48400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:763,&#34;encoded_body_size&#34;:420,&#34;decoded_body_size&#34;:763,&#34;transfer_size&#34;:720,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:48300000},&#34;first_byte&#34;:{&#34;duration&#34;:47900000,&#34;start&#34;:400000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078919,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6dff0525-975a-4a48-bdd3-b1da2195ca27&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCbef2e7ebac224955b821d9e70110c78b-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:48600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:921,&#34;encoded_body_size&#34;:512,&#34;decoded_body_size&#34;:921,&#34;transfer_size&#34;:812,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:48500000},&#34;first_byte&#34;:{&#34;duration&#34;:48000000,&#34;start&#34;:500000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;41f8c909-c4d0-4c76-b1fe-22a4b2d7f31d&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC0c3a7a5da52745d29b8ee369060b69c6-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:839,&#34;encoded_body_size&#34;:465,&#34;decoded_body_size&#34;:839,&#34;transfer_size&#34;:765,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:82000000},&#34;first_byte&#34;:{&#34;duration&#34;:80700000,&#34;start&#34;:1300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;20c217d3-c95a-4da5-a524-c54fbd48a63c&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC0e765aea33c34bfe885facaf6fc7febb-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1051,&#34;encoded_body_size&#34;:574,&#34;decoded_body_size&#34;:1051,&#34;transfer_size&#34;:874,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:82200000},&#34;first_byte&#34;:{&#34;duration&#34;:80900000,&#34;start&#34;:1300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1daee92a-2d65-49a3-a687-a90c532022d2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC30aa5696f5944a38a6fd12336395486c-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:432,&#34;encoded_body_size&#34;:275,&#34;decoded_body_size&#34;:432,&#34;transfer_size&#34;:575,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:82400000},&#34;first_byte&#34;:{&#34;duration&#34;:81100000,&#34;start&#34;:1300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d1330e90-d4ad-4706-a57e-85aa6645b6c8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC702605e1c6554864bc71e30c9d81d484-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:808,&#34;encoded_body_size&#34;:495,&#34;decoded_body_size&#34;:808,&#34;transfer_size&#34;:795,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:82400000},&#34;first_byte&#34;:{&#34;duration&#34;:81300000,&#34;start&#34;:1100000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6bfda98f-f23f-4b99-83bf-186e6ed07da2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://munchkin.marketo.net/munchkin.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:168900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b839e076-db1e-4de6-9208-64df8e549790&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCe9011213d2114262a74c6c798d020dc2-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:862,&#34;encoded_body_size&#34;:550,&#34;decoded_body_size&#34;:862,&#34;transfer_size&#34;:850,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:82400000},&#34;first_byte&#34;:{&#34;duration&#34;:81300000,&#34;start&#34;:1100000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;05868242-26d1-456d-b208-55eb41265f30&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC9147058850554edca24cff1971dc8752-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:984,&#34;encoded_body_size&#34;:590,&#34;decoded_body_size&#34;:984,&#34;transfer_size&#34;:890,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:82500000},&#34;first_byte&#34;:{&#34;duration&#34;:81500000,&#34;start&#34;:1000000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079059,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;2dc84e93-a5a5-4984-99e4-fe5d1c46929f&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCaf2d212b50ce48c1a6c3cd2be3fe5c2a-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:26900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1090,&#34;encoded_body_size&#34;:551,&#34;decoded_body_size&#34;:1090,&#34;transfer_size&#34;:851,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:26700000},&#34;first_byte&#34;:{&#34;duration&#34;:26400000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079102,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c4f7e951-5225-4111-9113-0c1041b3b7c8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCb6522d3edd5647029288904d3b3ea9e5-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:17200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:984,&#34;encoded_body_size&#34;:591,&#34;decoded_body_size&#34;:984,&#34;transfer_size&#34;:891,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:17100000},&#34;first_byte&#34;:{&#34;duration&#34;:16800000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079105,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;2e49864b-7164-48c7-814d-e5993a7f5214&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/ipv?_biz_r=&amp;_biz_h=-1719904874&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079103&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=0&amp;a=crowdstrike.com&amp;rnd=480761&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:18200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079105,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;295abb17-08bf-4820-94ad-ffdd9d463523&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizibly.com/u?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079104&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;a=crowdstrike.com&amp;rnd=110587&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:70500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079152,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8757d1ed-2395-44f4-8028-03f1cf9fac8d&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/xdc.js?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_h=-1719904874&amp;cdn_o=a&amp;jsVer=4.25.04.18&amp;a=crowdstrike.com&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:70800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079152,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;cbf483a2-442c-4675-bbae-c23316da2234&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC52d7b7dccc3548adaf8495e20b99dd8d-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:20700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1991,&#34;encoded_body_size&#34;:614,&#34;decoded_body_size&#34;:1991,&#34;transfer_size&#34;:914,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:20600000},&#34;first_byte&#34;:{&#34;duration&#34;:20300000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}`))
	

	
	req.Header.Set("Accept-Language", "en-US")
	
	req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != -1 {
		t.Errorf("Expected status %d, got %d", -1, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Axhr%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=8b72b0c9-5b91-4171-9212-cf72cfc10145&amp;amp;batch_time=1746729080202", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078919,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;a24cf75a-e1f8-4a89-b3b0-89f244b687cd&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/scripts/bizible.js?account=crowdstrike.com&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:182900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078919,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8438408b-dc0a-4c5c-8295-960f9e24bda2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCaf408a7d7bb64bba9c9ebd33e91577a5-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:48400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:763,&#34;encoded_body_size&#34;:420,&#34;decoded_body_size&#34;:763,&#34;transfer_size&#34;:720,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:48300000},&#34;first_byte&#34;:{&#34;duration&#34;:47900000,&#34;start&#34;:400000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078919,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6dff0525-975a-4a48-bdd3-b1da2195ca27&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCbef2e7ebac224955b821d9e70110c78b-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:48600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:921,&#34;encoded_body_size&#34;:512,&#34;decoded_body_size&#34;:921,&#34;transfer_size&#34;:812,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:48500000},&#34;first_byte&#34;:{&#34;duration&#34;:48000000,&#34;start&#34;:500000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;41f8c909-c4d0-4c76-b1fe-22a4b2d7f31d&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC0c3a7a5da52745d29b8ee369060b69c6-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:839,&#34;encoded_body_size&#34;:465,&#34;decoded_body_size&#34;:839,&#34;transfer_size&#34;:765,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:82000000},&#34;first_byte&#34;:{&#34;duration&#34;:80700000,&#34;start&#34;:1300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;20c217d3-c95a-4da5-a524-c54fbd48a63c&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC0e765aea33c34bfe885facaf6fc7febb-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1051,&#34;encoded_body_size&#34;:574,&#34;decoded_body_size&#34;:1051,&#34;transfer_size&#34;:874,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:82200000},&#34;first_byte&#34;:{&#34;duration&#34;:80900000,&#34;start&#34;:1300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1daee92a-2d65-49a3-a687-a90c532022d2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC30aa5696f5944a38a6fd12336395486c-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:432,&#34;encoded_body_size&#34;:275,&#34;decoded_body_size&#34;:432,&#34;transfer_size&#34;:575,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:82400000},&#34;first_byte&#34;:{&#34;duration&#34;:81100000,&#34;start&#34;:1300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d1330e90-d4ad-4706-a57e-85aa6645b6c8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC702605e1c6554864bc71e30c9d81d484-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:808,&#34;encoded_body_size&#34;:495,&#34;decoded_body_size&#34;:808,&#34;transfer_size&#34;:795,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:82400000},&#34;first_byte&#34;:{&#34;duration&#34;:81300000,&#34;start&#34;:1100000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6bfda98f-f23f-4b99-83bf-186e6ed07da2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://munchkin.marketo.net/munchkin.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:168900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b839e076-db1e-4de6-9208-64df8e549790&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCe9011213d2114262a74c6c798d020dc2-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:862,&#34;encoded_body_size&#34;:550,&#34;decoded_body_size&#34;:862,&#34;transfer_size&#34;:850,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:82400000},&#34;first_byte&#34;:{&#34;duration&#34;:81300000,&#34;start&#34;:1100000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078975,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;05868242-26d1-456d-b208-55eb41265f30&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC9147058850554edca24cff1971dc8752-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:82700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:984,&#34;encoded_body_size&#34;:590,&#34;decoded_body_size&#34;:984,&#34;transfer_size&#34;:890,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:82500000},&#34;first_byte&#34;:{&#34;duration&#34;:81500000,&#34;start&#34;:1000000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079059,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;2dc84e93-a5a5-4984-99e4-fe5d1c46929f&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCaf2d212b50ce48c1a6c3cd2be3fe5c2a-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:26900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1090,&#34;encoded_body_size&#34;:551,&#34;decoded_body_size&#34;:1090,&#34;transfer_size&#34;:851,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:26700000},&#34;first_byte&#34;:{&#34;duration&#34;:26400000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079102,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c4f7e951-5225-4111-9113-0c1041b3b7c8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCb6522d3edd5647029288904d3b3ea9e5-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:17200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:984,&#34;encoded_body_size&#34;:591,&#34;decoded_body_size&#34;:984,&#34;transfer_size&#34;:891,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:17100000},&#34;first_byte&#34;:{&#34;duration&#34;:16800000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079105,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;2e49864b-7164-48c7-814d-e5993a7f5214&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/ipv?_biz_r=&amp;_biz_h=-1719904874&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079103&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=0&amp;a=crowdstrike.com&amp;rnd=480761&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:18200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079105,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;295abb17-08bf-4820-94ad-ffdd9d463523&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizibly.com/u?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079104&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;a=crowdstrike.com&amp;rnd=110587&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:70500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079152,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8757d1ed-2395-44f4-8028-03f1cf9fac8d&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/xdc.js?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_h=-1719904874&amp;cdn_o=a&amp;jsVer=4.25.04.18&amp;a=crowdstrike.com&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:70800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079152,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;cbf483a2-442c-4675-bbae-c23316da2234&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC52d7b7dccc3548adaf8495e20b99dd8d-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:20700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1991,&#34;encoded_body_size&#34;:614,&#34;decoded_body_size&#34;:1991,&#34;transfer_size&#34;:914,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:20600000},&#34;first_byte&#34;:{&#34;duration&#34;:20300000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=a23ec7eb-27fe-4863-a867-e81decb82dce&amp;amp;batch_time=1746729080204", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079153,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;39c34b04-21e4-4cbe-9e95-00bd9f4d1b19&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://ad.doubleclick.net/activity;u1=www.crowdstrike.com%2Fen-us%2F;cat=crowd0;src=12037336;type=crowd0?&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:520900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:42,&#34;encoded_body_size&#34;:42,&#34;decoded_body_size&#34;:42,&#34;transfer_size&#34;:342,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:520900000},&#34;first_byte&#34;:{&#34;duration&#34;:61900000,&#34;start&#34;:459000000},&#34;connect&#34;:{&#34;duration&#34;:92200000,&#34;start&#34;:366100000},&#34;ssl&#34;:{&#34;duration&#34;:38500000,&#34;start&#34;:419800000},&#34;dns&#34;:{&#34;duration&#34;:1200000,&#34;start&#34;:364900000},&#34;redirect&#34;:{&#34;duration&#34;:362200000,&#34;start&#34;:0}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079155,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;babc1e7b-1f25-4ef5-aaab-e5331d7354f2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC429998db21a84d22962694934fe9c178-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:17400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:430,&#34;encoded_body_size&#34;:285,&#34;decoded_body_size&#34;:430,&#34;transfer_size&#34;:585,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:17200000},&#34;first_byte&#34;:{&#34;duration&#34;:17000000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079156,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;538441a1-00b3-48c0-8225-ed6058a5113b&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC3f830817bbac469aa1a50e321edd3939-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:16900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:690,&#34;encoded_body_size&#34;:425,&#34;decoded_body_size&#34;:690,&#34;transfer_size&#34;:725,&#34;download&#34;:{&#34;duration&#34;:900000,&#34;start&#34;:16000000},&#34;first_byte&#34;:{&#34;duration&#34;:15800000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079156,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5b26e21b-9362-41fa-9fc5-adc392ec9d46&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC4428e44be4d744ddac8973a220b18817-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:16500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:966,&#34;encoded_body_size&#34;:601,&#34;decoded_body_size&#34;:966,&#34;transfer_size&#34;:901,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:16300000},&#34;first_byte&#34;:{&#34;duration&#34;:16100000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079170,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;7499dcca-41ae-49dc-9724-1ed6f30f3e14&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://connect.facebook.net/signals/config/950083805267950?v=2.9.201&amp;r=stable&amp;domain=www.crowdstrike.com&amp;hme=9ebdfdd473ffce6bfe2267012c83f73483198ffe20d84139a2066b7682f827c0&amp;ex_m=73%2C128%2C113%2C117%2C64%2C6%2C106%2C72%2C19%2C100%2C92%2C54%2C57%2C181%2C202%2C209%2C205%2C206%2C208%2C32%2C107%2C56%2C80%2C207%2C176%2C179%2C203%2C204%2C189%2C139%2C44%2C194%2C191%2C192%2C37%2C151%2C18%2C53%2C198%2C197%2C141%2C21%2C43%2C2%2C46%2C68%2C69%2C70%2C74%2C96%2C20%2C17%2C99%2C95%2C94%2C114%2C55%2C116%2C42%2C115%2C33%2C97%2C29%2C177%2C180%2C148%2C14%2C15%2C16%2C8%2C9%2C28%2C25%2C26%2C60%2C65%2C67%2C78%2C105%2C108%2C30%2C79%2C12%2C10%2C83%2C51%2C24%2C110%2C109%2C111%2C102%2C13%2C23%2C4%2C41%2C77%2C22%2C160%2C89%2C135%2C76%2C1%2C98%2C59%2C87%2C36%2C31%2C85%2C86%2C91%2C40%2C7%2C93%2C84%2C47%2C35%2C38%2C0%2C71%2C118%2C90%2C5%2C50%2C49%2C101%2C88%2C246%2C174%2C126%2C163%2C156%2C3%2C39%2C66%2C45%2C112%2C48%2C82%2C63%2C62%2C34%2C103%2C61%2C58%2C52%2C81%2C75%2C27%2C104%2C11%2C119&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:24200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:7082,&#34;encoded_body_size&#34;:1271,&#34;decoded_body_size&#34;:7082,&#34;transfer_size&#34;:1571,&#34;download&#34;:{&#34;duration&#34;:2200000,&#34;start&#34;:22000000},&#34;first_byte&#34;:{&#34;duration&#34;:21500000,&#34;start&#34;:500000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079171,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b1eb40a6-33ad-425b-a7f8-787fe30a2997&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://munchkin.marketo.net/164/munchkin.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:23500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079188,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e44b991c-54f0-425c-a85c-7dc50f7eafbd&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://ob.fishrobotflower.com/i/771439ae128c64ffe20e624628cb6c78.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:109500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079188,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;12d899c7-4a27-42ef-a16e-ae041b2b95bc&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://trk.techtarget.com/tracking.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:132900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079189,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;513f1653-4b71-4eae-8c4c-28eb308b49f8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://sjrtp-cdn.marketo.com/rtp-api/v1/rtp.js?aid=crowdstrike&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:112000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079189,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c7876133-baf6-4e6e-8b1d-9fc3aba9ba15&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.redditstatic.com/ads/pixel.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:90300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079190,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6b9e032b-8a6c-4e4c-b06b-34ae31c35df2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://snap.licdn.com/li.lms-analytics/insight.min.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:83700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079190,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;95a02260-a5a2-4997-a7b0-d8a3fdfe5d29&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://www.influ2.com/tracker?clid=62c7557e-d1e3-40fb-93c4-d7c306706e53&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:104000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079191,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5523a31b-8969-4a51-b5b6-2ff81a11a8d3&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://ct.capterra.com/capterra_tracker.gif?vid=2104298&amp;vkey=884c38bc6ebbb2426278e18b331d9004&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:290300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079191,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;4703066e-8c4b-4429-952d-e8679dc46c55&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://arttrk.com/pixel/?ad_log=referer&amp;action=lead&amp;pixid=65acb722-b139-45a2-9a22-2e620e6d32b8&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:403400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079201,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;4c2927e4-a0f8-4910-865d-9dd0d2c180da&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC8a89dfd8b83c4511b4564c72e152a541-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:17200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:2311,&#34;encoded_body_size&#34;:942,&#34;decoded_body_size&#34;:2311,&#34;transfer_size&#34;:1242,&#34;download&#34;:{&#34;duration&#34;:1700000,&#34;start&#34;:15500000},&#34;first_byte&#34;:{&#34;duration&#34;:15300000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}`))
	

	
	req.Header.Set("Accept-Language", "en-US")
	
	req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != -1 {
		t.Errorf("Expected status %d, got %d", -1, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Axhr%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=21dc05c0-8089-4537-9268-98405e59d688&amp;amp;batch_time=1746729080204", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079153,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;39c34b04-21e4-4cbe-9e95-00bd9f4d1b19&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://ad.doubleclick.net/activity;u1=www.crowdstrike.com%2Fen-us%2F;cat=crowd0;src=12037336;type=crowd0?&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:520900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:42,&#34;encoded_body_size&#34;:42,&#34;decoded_body_size&#34;:42,&#34;transfer_size&#34;:342,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:520900000},&#34;first_byte&#34;:{&#34;duration&#34;:61900000,&#34;start&#34;:459000000},&#34;connect&#34;:{&#34;duration&#34;:92200000,&#34;start&#34;:366100000},&#34;ssl&#34;:{&#34;duration&#34;:38500000,&#34;start&#34;:419800000},&#34;dns&#34;:{&#34;duration&#34;:1200000,&#34;start&#34;:364900000},&#34;redirect&#34;:{&#34;duration&#34;:362200000,&#34;start&#34;:0}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079155,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;babc1e7b-1f25-4ef5-aaab-e5331d7354f2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC429998db21a84d22962694934fe9c178-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:17400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:430,&#34;encoded_body_size&#34;:285,&#34;decoded_body_size&#34;:430,&#34;transfer_size&#34;:585,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:17200000},&#34;first_byte&#34;:{&#34;duration&#34;:17000000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079156,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;538441a1-00b3-48c0-8225-ed6058a5113b&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC3f830817bbac469aa1a50e321edd3939-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:16900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:690,&#34;encoded_body_size&#34;:425,&#34;decoded_body_size&#34;:690,&#34;transfer_size&#34;:725,&#34;download&#34;:{&#34;duration&#34;:900000,&#34;start&#34;:16000000},&#34;first_byte&#34;:{&#34;duration&#34;:15800000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079156,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5b26e21b-9362-41fa-9fc5-adc392ec9d46&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC4428e44be4d744ddac8973a220b18817-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:16500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:966,&#34;encoded_body_size&#34;:601,&#34;decoded_body_size&#34;:966,&#34;transfer_size&#34;:901,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:16300000},&#34;first_byte&#34;:{&#34;duration&#34;:16100000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079170,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;7499dcca-41ae-49dc-9724-1ed6f30f3e14&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://connect.facebook.net/signals/config/950083805267950?v=2.9.201&amp;r=stable&amp;domain=www.crowdstrike.com&amp;hme=9ebdfdd473ffce6bfe2267012c83f73483198ffe20d84139a2066b7682f827c0&amp;ex_m=73%2C128%2C113%2C117%2C64%2C6%2C106%2C72%2C19%2C100%2C92%2C54%2C57%2C181%2C202%2C209%2C205%2C206%2C208%2C32%2C107%2C56%2C80%2C207%2C176%2C179%2C203%2C204%2C189%2C139%2C44%2C194%2C191%2C192%2C37%2C151%2C18%2C53%2C198%2C197%2C141%2C21%2C43%2C2%2C46%2C68%2C69%2C70%2C74%2C96%2C20%2C17%2C99%2C95%2C94%2C114%2C55%2C116%2C42%2C115%2C33%2C97%2C29%2C177%2C180%2C148%2C14%2C15%2C16%2C8%2C9%2C28%2C25%2C26%2C60%2C65%2C67%2C78%2C105%2C108%2C30%2C79%2C12%2C10%2C83%2C51%2C24%2C110%2C109%2C111%2C102%2C13%2C23%2C4%2C41%2C77%2C22%2C160%2C89%2C135%2C76%2C1%2C98%2C59%2C87%2C36%2C31%2C85%2C86%2C91%2C40%2C7%2C93%2C84%2C47%2C35%2C38%2C0%2C71%2C118%2C90%2C5%2C50%2C49%2C101%2C88%2C246%2C174%2C126%2C163%2C156%2C3%2C39%2C66%2C45%2C112%2C48%2C82%2C63%2C62%2C34%2C103%2C61%2C58%2C52%2C81%2C75%2C27%2C104%2C11%2C119&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:24200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:7082,&#34;encoded_body_size&#34;:1271,&#34;decoded_body_size&#34;:7082,&#34;transfer_size&#34;:1571,&#34;download&#34;:{&#34;duration&#34;:2200000,&#34;start&#34;:22000000},&#34;first_byte&#34;:{&#34;duration&#34;:21500000,&#34;start&#34;:500000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079171,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b1eb40a6-33ad-425b-a7f8-787fe30a2997&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://munchkin.marketo.net/164/munchkin.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:23500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079188,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e44b991c-54f0-425c-a85c-7dc50f7eafbd&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://ob.fishrobotflower.com/i/771439ae128c64ffe20e624628cb6c78.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:109500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079188,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;12d899c7-4a27-42ef-a16e-ae041b2b95bc&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://trk.techtarget.com/tracking.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:132900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079189,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;513f1653-4b71-4eae-8c4c-28eb308b49f8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://sjrtp-cdn.marketo.com/rtp-api/v1/rtp.js?aid=crowdstrike&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:112000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079189,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c7876133-baf6-4e6e-8b1d-9fc3aba9ba15&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.redditstatic.com/ads/pixel.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:90300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079190,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6b9e032b-8a6c-4e4c-b06b-34ae31c35df2&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://snap.licdn.com/li.lms-analytics/insight.min.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:83700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079190,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;95a02260-a5a2-4997-a7b0-d8a3fdfe5d29&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://www.influ2.com/tracker?clid=62c7557e-d1e3-40fb-93c4-d7c306706e53&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:104000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079191,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5523a31b-8969-4a51-b5b6-2ff81a11a8d3&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://ct.capterra.com/capterra_tracker.gif?vid=2104298&amp;vkey=884c38bc6ebbb2426278e18b331d9004&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:290300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079191,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;4703066e-8c4b-4429-952d-e8679dc46c55&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://arttrk.com/pixel/?ad_log=referer&amp;action=lead&amp;pixid=65acb722-b139-45a2-9a22-2e620e6d32b8&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:403400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079201,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;4c2927e4-a0f8-4910-865d-9dd0d2c180da&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC8a89dfd8b83c4511b4564c72e152a541-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:17200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:2311,&#34;encoded_body_size&#34;:942,&#34;decoded_body_size&#34;:2311,&#34;transfer_size&#34;:1242,&#34;download&#34;:{&#34;duration&#34;:1700000,&#34;start&#34;:15500000},&#34;first_byte&#34;:{&#34;duration&#34;:15300000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/u?mapType=ecid&amp;amp;mapValue=06D71E9261F941560A495CD6%40AdobeOrg_28180928486187528650392535887012291588&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729080106&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=2&amp;amp;a=crowdstrike.com&amp;amp;rnd=228061&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729080208", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://bat.bing.com/p/action/163002607.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://googleads.g.doubleclick.net/pagead/viewthroughconversion/797629828/?label=hozuCPn52LoYEIS7q_wC&amp;amp;guid=ON&amp;amp;script=0&amp;amp;ct_cookie_present=false&amp;amp;random=196357049&amp;amp;crd=CPLOsQIIorixAgihuLECCLHBsQIIsMGxAgixw7ECCIrFsQIIwsmxAgi0xrECCJDJsQIIw8-xAgjTxbECCOvMsQIIz86xAgj-zrECCNXPsQI&amp;amp;pscrd=IhMIof7j8sCUjQMVw5qOCB3oWQA-MgwIA2IICAAQABgAIAAyDAgEYggIABAAGAAgADIMCAdiCAgAEAAYACAAMgwICGIICAAQABgAIAAyDAgJYggIABAAGAAgADIMCApiCAgAEAAYACAAMgwIAmIICAAQABgAIAAyDAgLYggIABAAGAAgADIMCBViCAgAEAAYACAAMgwIH2IICAAQABgAIAAyDAgTYggIABAAGAAgADIMCBJiCAgAEAAYACAAOhxodHRwczovL3d3dy5jcm93ZHN0cmlrZS5jb20v", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/widgetapp/2025-05-08-12-13-34/locales/en-US.json", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://bat.bing.com/action/0?ti=163002607&amp;amp;Ver=2&amp;amp;mid=ae9bdae8-4702-4299-b1a6-34555b231a31&amp;amp;bo=1&amp;amp;sid=a2fec0b02c3a11f095697fbf91a88d6a&amp;amp;vid=a2fec0802c3a11f0b5c3b9fd0fc28d87&amp;amp;vids=1&amp;amp;msclkid=N&amp;amp;uach=pv%3D10.0&amp;amp;pi=0&amp;amp;lg=en-US&amp;amp;sw=1280&amp;amp;sh=720&amp;amp;sc=24&amp;amp;nwd=1&amp;amp;tl=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;p=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;r=&amp;amp;lt=569&amp;amp;evt=pageLoad&amp;amp;sv=1&amp;amp;cdb=AQET&amp;amp;rn=206258", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://bat.bing.com/action/0?ti=163002607&amp;amp;Ver=2&amp;amp;mid=ae9bdae8-4702-4299-b1a6-34555b231a31&amp;amp;bo=2&amp;amp;sid=a2fec0b02c3a11f095697fbf91a88d6a&amp;amp;vid=a2fec0802c3a11f0b5c3b9fd0fc28d87&amp;amp;vids=0&amp;amp;msclkid=N&amp;amp;ec=CHEQ&amp;amp;el=Invalid_Users&amp;amp;ev=0&amp;amp;ea=Invalid_Users&amp;amp;en=Y&amp;amp;p=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;sw=1280&amp;amp;sh=720&amp;amp;sc=24&amp;amp;nwd=1&amp;amp;evt=custom&amp;amp;cdb=AQET&amp;amp;rn=414587", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.google.com/pagead/1p-conversion/797629828/?label=hozuCPn52LoYEIS7q_wC&amp;amp;guid=ON&amp;amp;script=0&amp;amp;ct_cookie_present=false&amp;amp;random=196357049&amp;amp;crd=CPLOsQIIorixAgihuLECCLHBsQIIsMGxAgixw7ECCIrFsQIIwsmxAgi0xrECCJDJsQIIw8-xAgjTxbECCOvMsQIIz86xAgj-zrECCNXPsQI&amp;amp;pscrd=IhMIof7j8sCUjQMVw5qOCB3oWQA-MgwIA2IICAAQABgAIAAyDAgEYggIABAAGAAgADIMCAdiCAgAEAAYACAAMgwICGIICAAQABgAIAAyDAgJYggIABAAGAAgADIMCApiCAgAEAAYACAAMgwIAmIICAAQABgAIAAyDAgLYggIABAAGAAgADIMCBViCAgAEAAYACAAMgwIH2IICAAQABgAIAAyDAgTYggIABAAGAAgADIMCBJiCAgAEAAYACAAOhxodHRwczovL3d3dy5jcm93ZHN0cmlrZS5jb20v&amp;amp;is_vtc=1&amp;amp;cid=CAQSKQDZpuyzJReRKnL6vZeVEIjxxT7yvOWuKSXaMleGR-q-Y7BZ3eoUL74v&amp;amp;random=219317367", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=e808defb-2480-4f20-9d9c-7e2eefbd2036&amp;amp;batch_time=1746729080319", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079202,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;a5df2770-b8c2-4397-a3aa-9d0c00246ec7&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://connect.facebook.net/signals/config/992980065451679?v=2.9.201&amp;r=stable&amp;domain=www.crowdstrike.com&amp;hme=9ebdfdd473ffce6bfe2267012c83f73483198ffe20d84139a2066b7682f827c0&amp;ex_m=73%2C128%2C113%2C117%2C64%2C6%2C106%2C72%2C19%2C100%2C92%2C54%2C57%2C181%2C202%2C209%2C205%2C206%2C208%2C32%2C107%2C56%2C80%2C207%2C176%2C179%2C203%2C204%2C189%2C139%2C44%2C194%2C191%2C192%2C37%2C151%2C18%2C53%2C198%2C197%2C141%2C21%2C43%2C2%2C46%2C68%2C69%2C70%2C74%2C96%2C20%2C17%2C99%2C95%2C94%2C114%2C55%2C116%2C42%2C115%2C33%2C97%2C29%2C177%2C180%2C148%2C14%2C15%2C16%2C8%2C9%2C28%2C25%2C26%2C60%2C65%2C67%2C78%2C105%2C108%2C30%2C79%2C12%2C10%2C83%2C51%2C24%2C110%2C109%2C111%2C102%2C13%2C23%2C4%2C41%2C77%2C22%2C160%2C89%2C135%2C76%2C1%2C98%2C59%2C87%2C36%2C31%2C85%2C86%2C91%2C40%2C7%2C93%2C84%2C47%2C35%2C38%2C0%2C71%2C118%2C90%2C5%2C50%2C49%2C101%2C88%2C246%2C174%2C126%2C163%2C156%2C3%2C39%2C66%2C45%2C112%2C48%2C82%2C63%2C62%2C34%2C103%2C61%2C58%2C52%2C81%2C75%2C27%2C104%2C11%2C119&#34;,&#34;protocol&#34;:&#34;h3&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:142500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:68543,&#34;encoded_body_size&#34;:13852,&#34;decoded_body_size&#34;:68543,&#34;transfer_size&#34;:14152,&#34;download&#34;:{&#34;duration&#34;:300000,&#34;start&#34;:142200000},&#34;first_byte&#34;:{&#34;duration&#34;:142000000,&#34;start&#34;:200000},&#34;connect&#34;:{&#34;duration&#34;:0,&#34;start&#34;:100000},&#34;ssl&#34;:{&#34;duration&#34;:0,&#34;start&#34;:100000},&#34;dns&#34;:{&#34;duration&#34;:0,&#34;start&#34;:100000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079202,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8a0d51ea-8eb8-4a8e-b1f4-9a1386646698&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.facebook.com/tr/?id=950083805267950&amp;ev=PageView&amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;rl=&amp;if=false&amp;ts=1746729079201&amp;sw=1280&amp;sh=720&amp;v=2.9.201&amp;r=stable&amp;a=adobe_launch&amp;ec=0&amp;o=28&amp;it=1746729079168&amp;coo=false&amp;exp=k0&amp;rqm=GET&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:129700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079205,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;3099c754-99f5-4c92-8869-0b9e76ec50a7&#34;,&#34;type&#34;:&#34;beacon&#34;,&#34;url&#34;:&#34;https://281-obq-266.mktoresp.com/webevents/visitWebPage?_mchNc=1746729079204&amp;_mchCn=&amp;_mchId=281-OBQ-266&amp;_mchTk=_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;_mchHo=www.crowdstrike.com&amp;_mchPo=&amp;_mchRu=%2Fen-us%2F&amp;_mchPc=https%3A&amp;_mchVr=164&amp;aip=1&amp;_mchEcid=06D71E9261F941560A495CD6%40AdobeOrg%3A%3A28180928486187528650392535887012291588&amp;_mchHa=&amp;_mchRe=&amp;_mchQp=&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:435900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079473,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;56194d70-b37a-45b4-9b41-bd9907c9ca8c&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCe54f772526294e1696f6be4e42252bd6-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:17200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:645,&#34;encoded_body_size&#34;:400,&#34;decoded_body_size&#34;:645,&#34;transfer_size&#34;:700,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:17200000},&#34;first_byte&#34;:{&#34;duration&#34;:16700000,&#34;start&#34;:500000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079474,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;91350a31-dd3d-4e89-999a-4079cc50f621&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://ad.doubleclick.net/activity;u1=www.crowdstrike.com%2Fen-us%2F;cat=crowd00;src=12037336;type=crowd0?&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:199100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:42,&#34;encoded_body_size&#34;:42,&#34;decoded_body_size&#34;:42,&#34;transfer_size&#34;:342,&#34;download&#34;:{&#34;duration&#34;:600000,&#34;start&#34;:198500000},&#34;first_byte&#34;:{&#34;duration&#34;:60500000,&#34;start&#34;:138000000},&#34;redirect&#34;:{&#34;duration&#34;:135400000,&#34;start&#34;:0}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079474,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b27de44a-a717-4a6d-a596-002e23ed01d9&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC13bb3acaad764c19a86cb66a819d9b01-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:16500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:551,&#34;encoded_body_size&#34;:328,&#34;decoded_body_size&#34;:551,&#34;transfer_size&#34;:628,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:16300000},&#34;first_byte&#34;:{&#34;duration&#34;:16200000,&#34;start&#34;:100000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079475,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;4e6aed2e-0a41-439b-94ba-4847949571d8&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://t.contentsquare.net/uxa/184b355acd0d7.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:133600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:465773,&#34;encoded_body_size&#34;:105554,&#34;decoded_body_size&#34;:465773,&#34;transfer_size&#34;:105854,&#34;download&#34;:{&#34;duration&#34;:34000000,&#34;start&#34;:99600000},&#34;first_byte&#34;:{&#34;duration&#34;:62400000,&#34;start&#34;:37200000},&#34;connect&#34;:{&#34;duration&#34;:35700000,&#34;start&#34;:1400000},&#34;ssl&#34;:{&#34;duration&#34;:19200000,&#34;start&#34;:17900000},&#34;dns&#34;:{&#34;duration&#34;:1100000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079476,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b9015ede-d7b7-43d6-b9ee-f5c7a3017360&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC8a19ad85a7d044158f34540c69c7d8b4-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:22400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:3615,&#34;encoded_body_size&#34;:1154,&#34;decoded_body_size&#34;:3615,&#34;transfer_size&#34;:1454,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:22300000},&#34;first_byte&#34;:{&#34;duration&#34;:22000000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079476,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;abcd86d8-6b33-40ee-be4c-f650e60beb51&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCca072a9a5f2f46d19d4d38bdcb9d6873-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:22500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1987,&#34;encoded_body_size&#34;:1049,&#34;decoded_body_size&#34;:1987,&#34;transfer_size&#34;:1349,&#34;download&#34;:{&#34;duration&#34;:300000,&#34;start&#34;:22200000},&#34;first_byte&#34;:{&#34;duration&#34;:21900000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079488,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d2d57143-2324-40a6-b5fe-99909be35ada&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://alb.reddit.com/rp.gif?ts=1746729079487&amp;id=t2_2n40s6z5&amp;event=PageVisit&amp;m.itemCount=&amp;m.value=&amp;m.valueDecimal=&amp;m.currency=&amp;m.transactionId=&amp;m.customEventName=&amp;m.products=&amp;m.conversionId=&amp;uuid=2bc44b75-14f8-4772-a1c1-5d39c30b5042&amp;aaid=&amp;em=&amp;pn=&amp;external_id=&amp;idfa=&amp;integration=reddit&amp;partner=&amp;opt_out=0&amp;sh=1280&amp;sw=720&amp;v=rdt_00917691&amp;dpm=&amp;dpcc=&amp;dprc=&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:242000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079498,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;88903d89-8334-4cd3-a53a-8c37d19ff091&#34;,&#34;type&#34;:&#34;beacon&#34;,&#34;url&#34;:&#34;https://www.google.com/ccm/collect?tid=AW-797629828&amp;en=page_view&amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;scrsrc=www.googletagmanager.com&amp;frm=0&amp;rnd=2104649365.1746729079&amp;dt=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;auid=1475816545.1746729079&amp;navt=n&amp;npa=0&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;tft=1746729079497&amp;tfd=1363&amp;apve=1&amp;apvf=sb&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:169000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079502,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5dc4c4ea-2801-45cc-b555-03081e3fc92b&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://googleads.g.doubleclick.net/pagead/viewthroughconversion/797629828/?random=1746729079496&amp;cv=11&amp;fst=1746729079496&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&amp;rfmt=3&amp;fmt=4&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:185600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:3981,&#34;encoded_body_size&#34;:1996,&#34;decoded_body_size&#34;:3981,&#34;transfer_size&#34;:2296,&#34;download&#34;:{&#34;duration&#34;:800000,&#34;start&#34;:184800000},&#34;first_byte&#34;:{&#34;duration&#34;:75600000,&#34;start&#34;:109200000},&#34;connect&#34;:{&#34;duration&#34;:76400000,&#34;start&#34;:32000000},&#34;ssl&#34;:{&#34;duration&#34;:38500000,&#34;start&#34;:69900000},&#34;dns&#34;:{&#34;duration&#34;:31800000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=290eba1b-c6b1-453e-a409-60bfd7d28347&amp;amp;batch_time=1746729080341", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079503,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8b88237a-fa75-467b-bc47-891b693017cf&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://td.doubleclick.net/td/rul/797629828?random=1746729079496&amp;cv=11&amp;fst=1746729079496&amp;fmt=3&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:205300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:3750,&#34;encoded_body_size&#34;:1161,&#34;decoded_body_size&#34;:3750,&#34;transfer_size&#34;:1461,&#34;download&#34;:{&#34;duration&#34;:300000,&#34;start&#34;:205000000},&#34;first_byte&#34;:{&#34;duration&#34;:96600000,&#34;start&#34;:108400000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079505,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;70b37435-db8d-44ee-ab8d-c21b15e429c4&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://googleads.g.doubleclick.net/pagead/viewthroughconversion/952416460/?random=1746729079503&amp;cv=11&amp;fst=1746729079503&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&amp;rfmt=3&amp;fmt=4&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:183100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:3982,&#34;encoded_body_size&#34;:1996,&#34;decoded_body_size&#34;:3982,&#34;transfer_size&#34;:2296,&#34;download&#34;:{&#34;duration&#34;:900000,&#34;start&#34;:182200000},&#34;first_byte&#34;:{&#34;duration&#34;:75800000,&#34;start&#34;:106400000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079506,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;317284b0-323c-41b6-8140-b67a6912e489&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://td.doubleclick.net/td/rul/952416460?random=1746729079503&amp;cv=11&amp;fst=1746729079503&amp;fmt=3&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:186300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:2267,&#34;encoded_body_size&#34;:941,&#34;decoded_body_size&#34;:2267,&#34;transfer_size&#34;:1241,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:186100000},&#34;first_byte&#34;:{&#34;duration&#34;:80300000,&#34;start&#34;:105800000},&#34;connect&#34;:{&#34;duration&#34;:76700000,&#34;start&#34;:28500000},&#34;ssl&#34;:{&#34;duration&#34;:38400000,&#34;start&#34;:66800000},&#34;dns&#34;:{&#34;duration&#34;:28200000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079506,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c7736b4f-e5f9-48da-8da2-8c8f68b7130d&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.datadoghq-browser-agent.com/us1/v5/datadog-rum.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:240800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:165703,&#34;encoded_body_size&#34;:54841,&#34;decoded_body_size&#34;:165703,&#34;transfer_size&#34;:55141,&#34;download&#34;:{&#34;duration&#34;:22000000,&#34;start&#34;:218800000},&#34;first_byte&#34;:{&#34;duration&#34;:52900000,&#34;start&#34;:165900000},&#34;connect&#34;:{&#34;duration&#34;:42100000,&#34;start&#34;:123700000},&#34;ssl&#34;:{&#34;duration&#34;:24200000,&#34;start&#34;:141600000},&#34;dns&#34;:{&#34;duration&#34;:35200000,&#34;start&#34;:88500000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079515,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;ffef9145-b137-4640-9559-439418a7d6a9&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://rtp-static.marketo.com/rtp/libs/jquery-ui-insightera-custom-1.9.6.css&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:152800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079516,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c8716810-40f3-4578-b72c-8a1163b3ff2c&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://rtp-static.marketo.com/rtp/libs/ga-integration-2.0.5.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:211700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079524,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6c561e2e-2287-4457-8303-19568bb09d6c&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.facebook.com/tr/?id=992980065451679&amp;ev=PageView&amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;rl=&amp;if=false&amp;ts=1746729079523&amp;sw=1280&amp;sh=720&amp;v=2.9.201&amp;r=stable&amp;ec=0&amp;o=4126&amp;fbp=fb.1.1746729079520.75795469452619173&amp;ler=empty&amp;it=1746729079168&amp;coo=false&amp;exp=k2&amp;rqm=GET&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:51500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079717,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;aaf6afc0-7ce9-4882-9208-d44421e85325&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.google.com/pagead/1p-user-list/797629828/?random=1746729079496&amp;cv=11&amp;fst=1746727200000&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&amp;rfmt=3&amp;fmt=3&amp;is_vtc=1&amp;cid=CAQSGwDZpuyz7viBNmNM2I-21Pt81hd1C8rcQz-rtg&amp;random=246964701&amp;rmt_tld=0&amp;ipr=y&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:47600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:42,&#34;encoded_body_size&#34;:42,&#34;decoded_body_size&#34;:42,&#34;transfer_size&#34;:342,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:47600000},&#34;first_byte&#34;:{&#34;duration&#34;:47000000,&#34;start&#34;:600000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079717,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;acaead79-19e6-4d50-8852-bace6748a246&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.google.com/pagead/1p-user-list/952416460/?random=1746729079503&amp;cv=11&amp;fst=1746727200000&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&amp;rfmt=3&amp;fmt=3&amp;is_vtc=1&amp;cid=CAQSGwDZpuyzPChu2ZLBUSUbxC76xLhpZO1EVFB39g&amp;random=3015211708&amp;rmt_tld=0&amp;ipr=y&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:58500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:42,&#34;encoded_body_size&#34;:42,&#34;decoded_body_size&#34;:42,&#34;transfer_size&#34;:342,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:58400000},&#34;first_byte&#34;:{&#34;duration&#34;:58000000,&#34;start&#34;:400000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079515,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e291890d-b170-4a6a-b9db-16e1fcb883ae&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://rtp-static.marketo.com/rtp/libs/jquery/3.7.0/jquery.min.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:260800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079669,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;37ae88af-cf3b-4716-8bcb-67c05bbc7e50&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://ext.chtbl.com/trackable.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:193100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079516,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;46077a65-614a-44b4-8364-7918b8c0307e&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/trw?aid=crowdstrike&amp;trwv.uid=crowdstrike-1746729079514-1f81ceec&amp;trwv.vc=1&amp;trwsa.sid=crowdstrike-1746729079515-66d02f45&amp;trwsb.cpv=1&amp;ctzo=-07:00&amp;uri=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;pm=&amp;viewedTypes=&amp;rts=1746729079515&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:346400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079668,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5b1fbd71-b584-43be-ad74-56bfd2158b01&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/msg?a=2&amp;sid=crowdstrike-1746729079515-66d02f45&amp;aid=crowdstrike&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;viewedTypes=&amp;0.08197428152563013&amp;rts=1746729079668&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:211600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	
	req.Header.Set("Accept-Language", "en-US")
	
	req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != -1 {
		t.Errorf("Expected status %d, got %d", -1, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Axhr%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=64db1b83-3ca6-4da4-846a-a0c41497f4a0&amp;amp;batch_time=1746729080341", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079503,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8b88237a-fa75-467b-bc47-891b693017cf&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://td.doubleclick.net/td/rul/797629828?random=1746729079496&amp;cv=11&amp;fst=1746729079496&amp;fmt=3&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:205300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:3750,&#34;encoded_body_size&#34;:1161,&#34;decoded_body_size&#34;:3750,&#34;transfer_size&#34;:1461,&#34;download&#34;:{&#34;duration&#34;:300000,&#34;start&#34;:205000000},&#34;first_byte&#34;:{&#34;duration&#34;:96600000,&#34;start&#34;:108400000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079505,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;70b37435-db8d-44ee-ab8d-c21b15e429c4&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://googleads.g.doubleclick.net/pagead/viewthroughconversion/952416460/?random=1746729079503&amp;cv=11&amp;fst=1746729079503&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&amp;rfmt=3&amp;fmt=4&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:183100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:3982,&#34;encoded_body_size&#34;:1996,&#34;decoded_body_size&#34;:3982,&#34;transfer_size&#34;:2296,&#34;download&#34;:{&#34;duration&#34;:900000,&#34;start&#34;:182200000},&#34;first_byte&#34;:{&#34;duration&#34;:75800000,&#34;start&#34;:106400000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079506,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;317284b0-323c-41b6-8140-b67a6912e489&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://td.doubleclick.net/td/rul/952416460?random=1746729079503&amp;cv=11&amp;fst=1746729079503&amp;fmt=3&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:186300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:2267,&#34;encoded_body_size&#34;:941,&#34;decoded_body_size&#34;:2267,&#34;transfer_size&#34;:1241,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:186100000},&#34;first_byte&#34;:{&#34;duration&#34;:80300000,&#34;start&#34;:105800000},&#34;connect&#34;:{&#34;duration&#34;:76700000,&#34;start&#34;:28500000},&#34;ssl&#34;:{&#34;duration&#34;:38400000,&#34;start&#34;:66800000},&#34;dns&#34;:{&#34;duration&#34;:28200000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079506,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c7736b4f-e5f9-48da-8da2-8c8f68b7130d&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://www.datadoghq-browser-agent.com/us1/v5/datadog-rum.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:240800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:165703,&#34;encoded_body_size&#34;:54841,&#34;decoded_body_size&#34;:165703,&#34;transfer_size&#34;:55141,&#34;download&#34;:{&#34;duration&#34;:22000000,&#34;start&#34;:218800000},&#34;first_byte&#34;:{&#34;duration&#34;:52900000,&#34;start&#34;:165900000},&#34;connect&#34;:{&#34;duration&#34;:42100000,&#34;start&#34;:123700000},&#34;ssl&#34;:{&#34;duration&#34;:24200000,&#34;start&#34;:141600000},&#34;dns&#34;:{&#34;duration&#34;:35200000,&#34;start&#34;:88500000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079515,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;ffef9145-b137-4640-9559-439418a7d6a9&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://rtp-static.marketo.com/rtp/libs/jquery-ui-insightera-custom-1.9.6.css&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:152800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079516,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c8716810-40f3-4578-b72c-8a1163b3ff2c&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://rtp-static.marketo.com/rtp/libs/ga-integration-2.0.5.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:211700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079524,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6c561e2e-2287-4457-8303-19568bb09d6c&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.facebook.com/tr/?id=992980065451679&amp;ev=PageView&amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;rl=&amp;if=false&amp;ts=1746729079523&amp;sw=1280&amp;sh=720&amp;v=2.9.201&amp;r=stable&amp;ec=0&amp;o=4126&amp;fbp=fb.1.1746729079520.75795469452619173&amp;ler=empty&amp;it=1746729079168&amp;coo=false&amp;exp=k2&amp;rqm=GET&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:51500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079717,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;aaf6afc0-7ce9-4882-9208-d44421e85325&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.google.com/pagead/1p-user-list/797629828/?random=1746729079496&amp;cv=11&amp;fst=1746727200000&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&amp;rfmt=3&amp;fmt=3&amp;is_vtc=1&amp;cid=CAQSGwDZpuyz7viBNmNM2I-21Pt81hd1C8rcQz-rtg&amp;random=246964701&amp;rmt_tld=0&amp;ipr=y&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:47600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:42,&#34;encoded_body_size&#34;:42,&#34;decoded_body_size&#34;:42,&#34;transfer_size&#34;:342,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:47600000},&#34;first_byte&#34;:{&#34;duration&#34;:47000000,&#34;start&#34;:600000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079717,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;acaead79-19e6-4d50-8852-bace6748a246&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.google.com/pagead/1p-user-list/952416460/?random=1746729079503&amp;cv=11&amp;fst=1746727200000&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;data=event%3Dgtag.config&amp;rfmt=3&amp;fmt=3&amp;is_vtc=1&amp;cid=CAQSGwDZpuyzPChu2ZLBUSUbxC76xLhpZO1EVFB39g&amp;random=3015211708&amp;rmt_tld=0&amp;ipr=y&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:58500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:42,&#34;encoded_body_size&#34;:42,&#34;decoded_body_size&#34;:42,&#34;transfer_size&#34;:342,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:58400000},&#34;first_byte&#34;:{&#34;duration&#34;:58000000,&#34;start&#34;:400000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079515,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e291890d-b170-4a6a-b9db-16e1fcb883ae&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://rtp-static.marketo.com/rtp/libs/jquery/3.7.0/jquery.min.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:260800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079669,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;37ae88af-cf3b-4716-8bcb-67c05bbc7e50&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://ext.chtbl.com/trackable.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:193100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079516,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;46077a65-614a-44b4-8364-7918b8c0307e&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/trw?aid=crowdstrike&amp;trwv.uid=crowdstrike-1746729079514-1f81ceec&amp;trwv.vc=1&amp;trwsa.sid=crowdstrike-1746729079515-66d02f45&amp;trwsb.cpv=1&amp;ctzo=-07:00&amp;uri=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;pm=&amp;viewedTypes=&amp;rts=1746729079515&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:346400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079668,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5b1fbd71-b584-43be-ad74-56bfd2158b01&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/msg?a=2&amp;sid=crowdstrike-1746729079515-66d02f45&amp;aid=crowdstrike&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;viewedTypes=&amp;0.08197428152563013&amp;rts=1746729079668&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:211600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://googleads.g.doubleclick.net/pagead/viewthroughconversion/797629828/?random=160478099&amp;amp;cv=11&amp;amp;fst=1746729080069&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gcl_ctr=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;label=hozuCPn52LoYEIS7q_wC&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;gtm_ee=1&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;capi=1&amp;amp;data=event%3Dconversion&amp;amp;fmt=3&amp;amp;ct_cookie_present=false&amp;amp;crd=CPLOsQIIobixAgixwbECCLDBsQIIscOxAgiKxbECCMLJsQIItMaxAgiQybECCNPFsQII68yxAgjPzrECCP7OsQII1c-xAkoVdHJpZ2dlciwgZXZlbnQtc291cmNlWgMKAQFiBAoCAgM&amp;amp;pscrd=IhMI34Lk8sCUjQMVqKOOCB2_0BUQMgwIA2IICAAQABgAIAAyDAgEYggIABAAGAAgADIMCAdiCAgAEAAYACAAMgwICGIICAAQABgAIAAyDAgJYggIABAAGAAgADIMCApiCAgAEAAYACAAMgwIAmIICAAQABgAIAAyDAgLYggIABAAGAAgADIMCBViCAgAEAAYACAAMgwIH2IICAAQABgAIAAyDAgTYggIABAAGAAgADIMCBJiCAgAEAAYACAAOhxodHRwczovL3d3dy5jcm93ZHN0cmlrZS5jb20vQlZDaEVJOEt6eHdBWVFvTE9vdXAtQXU5bjlBUklyQVBSSTZXMURUUl94MUZMVUdqZUNFdGtHbWhGY0RkQmdnRHdlMVQ4SUNUYVpZRTdHN095bFlqY0ZBUQ", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		t.Errorf("Expected status %d, got %d", 302, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=c814a790-1591-4a7a-a1a2-4c64725fc14a&amp;amp;batch_time=1746729080416", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079645,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;90067582-4fd3-4277-b9a6-028d1928ebeb&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.userway.org/widget.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:259500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079781,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e4efc182-2571-49eb-b8ea-de1da9120daa&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://rtp-static.marketo.com/rtp/libs/jqueryui/1.13.2/jquery-custom-ui.min.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:129500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079189,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;891b5f2d-3ae8-4705-8b8a-49518d394e20&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://collector-20290.tvsquared.com/tv2track.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:809800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079865,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c11eadd8-7ac9-4cc9-9deb-c6510131af66&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/msg?a=2&amp;sid=crowdstrike-1746729079515-66d02f45&amp;aid=crowdstrike&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;viewedTypes=&amp;0.5769826221088862&amp;rts=1746729079864&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:133900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079485,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;e827e592-2a3d-440d-b6b1-a68c6598a43f&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://px.ads.linkedin.com/collect?v=2&amp;fmt=js&amp;pid=64444&amp;time=1746729079483&amp;li_adsId=f8924343-8fd8-45cc-9a32-420fa0932ae3&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:571800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079648,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;de57e4ed-a381-46c9-a1cf-f6968e33155d&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://obs.fishrobotflower.com/ct?id=42110&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;sf=0&amp;tpi=&amp;ch=cheq4ppc&amp;uvid=undefined&amp;tsf=0&amp;tsfmi=&amp;tsfu=&amp;cb=1746729079646&amp;hl=2&amp;op=0&amp;ag=1932998597&amp;rand=8452781980417220858212893201002921909569170169186085261562527102269915086580215220293222&amp;fs=1280x720&amp;fst=1280x720&amp;np=macintel&amp;nv=google%20inc.&amp;ref=&amp;ss=1280x720&amp;nc=0&amp;at=&amp;di=W1siZWYiLDQwNTVdLFsiYWJuY2giLDVdLFstMTMsIi0iXSxbLTIyLCJbXCJuXCIsXCJuXCJdIl0sWy0zOCwiaSwtMSwtMSwwLDAsMCwwLDIsNTYsMTAxLC0xLDAsNTE2LDUxNiwxMzk1LDEzOTYiXSxbLTMsIltdIl0sWy0xMCwiLSJdLFstNDIsIjg4MzM5OTAxNiJdLFstNDUsIjYyMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsNjMwLDU3MSw0OTMsNjA4LDc5Miw2NzcsMCwwLDAsMCwwIl0sWzEyLCJ7XCJjdHhcIjpcIndlYmdsXCIsXCJ2XCI6XCJnb29nbGUgaW5jLiAoYXBwbGUpXCIsXCJyXCI6XCJhbmdsZSAoYXBwbGUsIGFwcGxlIG00IHBybywgb3BlbmdsIDQuMSlcIixcInNsdlwiOlwid2ViZ2wgZ2xzbCBlcyAxLjAgKG9wZW5nbCBlcyBnbHNsIGVzIDEuMCBjaHJvbWl1bSlcIixcImd2ZXJcIjpcIndlYmdsIDEuMCAob3BlbmdsIGVzIDIuMCBjaHJvbWl1bSlcIixcImd2ZW5cIjpcIndlYmtpdFwiLFwiYmVuXCI6MixcIndnbFwiOjEsXCJncmVuXCI6XCJ3ZWJraXQgd2ViZ2xcIixcInNlZlwiOjI5NzA0MTcwOCxcInNlY1wiOlwiXCJ9Il0sWy0xOSwiWzAsMCwwLDAsMCwwLDEsMjQsMjQsXCItXCIsMTI4MCw3MjAsMTI4MCw3MjAsMTI4MCw3MjAsMTI4MCw3MjAsMCwwLDAsMCxcIi1cIixcIi1cIiwxMjgwLDcyMCwwXSJdLFstMjYsIntcInRqaHNcIjozNTEwMDAwMCxcInVqaHNcIjoyNjAwMDAwMCxcImpoc2xcIjozNzYwMDAwMDAwfSJdLFstMjcsIlswLDEwLDAsXCI0Z1wiLG51bGxdIl0sWy0zNywiLSJdLFstNTUsIjEiXSxbLTUyLCItIl0sWy02MSwie1wid2dzbFwiOlwiNDtwYWNrZWRfNHg4X2ludGVnZXJfZG90X3Byb2R1Y3Q7dW5yZXN0cmljdGVkX3BvaW50ZXJfcGFyYW1ldGVycztwb2ludGVyX2NvbXBvc2l0ZV9hY2Nlc3M7cmVhZG9ubHlfYW5kX3JlYWR3cml0ZV9zdG9yYWdlX3RleHR1cmVzO1wiLFwicGNmXCI6XCJiZ3JhOHVub3JtXCJ9Il0sWy02NCwiWzAsXCJXaW5kb3dzXCIsW3tcImJcIjpcIkNocm9taXVtXCIsXCJ2XCI6XCIxMzZcIn0se1wiYlwiOlwiSGVhZGxlc3NDaHJvbWVcIixcInZcIjpcIjEzNlwifSx7XCJiXCI6XCJOb3QuQS9CcmFuZFwiLFwidlwiOlwiOTlcIn1dXSJdLFstMTQsIi0iXSxbLTIxLCItIl0sWy0yMywiKyJdLFstMjQsIltcInNheXN3aG9cIiwwLFwiQ2hyb21lIDEzNlwiLDEsMV0iXSxbLTQwLCIzMyJdLFstNTYsImxhbmRzY2FwZS1wcmltYXJ5Il0sWy01LCItIl0sWy02LCJ7XCJ3XCI6W1wiMFwiLFwiT25lVHJ1c3RTdHViXCIsXCJPcHRhbm9uV3JhcHBlclwiLFwiJFwiLFwialF1ZXJ5XCIsXCJtYXRjaGVkXCIsXCJicm93c2VyXCIsXCJHcmFuaXRlXCIsXCJFbmxpZ2h0ZXJKU1wiLFwidmlkeWFyZEVtYmVkXCIsXCJzZXRJbW1lZGlhdGVcIixcImNsZWFySW1tZWRpYXRlXCIsXCJWaWR5YXJkVjRcIixcIlZpZHlhcmRcIixcIk90VHJ1c3RlZFR5cGVcIixcIl9fU1ZHX1NQUklURV9fXCIsXCJBZGRTZWFyY2hDbGllbnRcIixcIkFkZFNlYXJjaFVJXCIsXCJnc2FwVmVyc2lvbnNcIixcIkNRXCIsXCJfc2xpY2VkVG9BcnJheVwiLFwiX25vbkl0ZXJhYmxlUmVzdFwiLFwiX2l0ZXJhYmxlVG9BcnJheUxpbWl0XCIsXCJfYXJyYXlXaXRoSG9sZXNcIixcIl9jcmVhdGVGb3JPZkl0ZXJhdG9ySGVscGVyXCIsXCJfdW5zdXBwb3J0ZWRJdGVyYWJsZVRvQXJyYXlcIixcIl9hcnJheUxpa2VUb0FycmF5XCIsXCJfdHlwZW9mXCIsXCJDTVBcIixcImFkb2JlRGF0YUxheWVyXCIsXCJyZWFjdGl2ZUVsZW1lbnRWZXJzaW9uc1wiLFwibGl0SHRtbFZlcnNpb25zXCIsXCJsaXRFbGVtZW50VmVyc2lvbnNcIixcIkNTU1J1bGVQbHVnaW5cIixcIkN1c3RvbUVhc2VcIixcIkRyYXdTVkdQbHVnaW5cIixcIkVhc2VsUGx1Z2luXCIsXCJFYXNlUGFja1wiLFwiRXhwb1NjYWxlRWFzZVwiLFwiUm91Z2hFYXNlXCIsXCJTbG93TW9cIixcIkxpbmVhclwiLFwiUG93ZXIwXCIsXCJRdWFkXCIsXCJQb3dlcjFcIixcIkN1YmljXCIsXCJQb3dlcjJcIixcIlF1YXJ0XCIsXCJQb3dlcjNcIixcIlF1aW50XCJdLFwiblwiOltcInNheXN3aG9cIl0sXCJkXCI6W119Il0sWy05LCIrIl0sWy0xNywiMTIiXSxbLTIwLCI0NjIyNTAyNDEuMTc0NjcyOTA3OSJdLFstMjUsIi0iXSxbLTM0LCItIl0sWy0zNiwiW1wiMTYvOVwiLFwiMTYvOVwiXSJdLFstNDMsIjAwMDAwMDAxMDAwMDAwMDAxMDExMTAxMTAwMTExMTAxMDAwMDAxMCJdLFstNDksIi0iXSxbLTUzLCIwMDEiXSxbLTU4LCItIl0sWy02MCwiLSJdLFstNjIsIjgwIl0sWy0xLCItIl0sWy03LCItIl0sWy0xMSwie1widFwiOlwiXCIsXCJtXCI6W1wiZGVzY3JpcHRpb25cIixcIm9nOnRpdGxlXCIsXCJvZzpkZXNjcmlwdGlvblwiLFwidHdpdHRlcjp0aXRsZVwiLFwidHdpdHRlcjpkZXNjcmlwdGlvblwiXX0iXSxbLTMwLCJbXCJ2XCIsMF0iXSxbLTMzLCItIl0sWy03MSwiYTAxMDAxMDExMDAxMDAxMDEwMDAxMDEwMDExMTExMDEwMDAwMTAiXSxbLTQsIi0iXSxbLTgsIi0iXSxbLTEyLCJudWxsIl0sWy0xNiwiMCJdLFstMTgsIlswLDAsMCwxXSJdLFstMzIsIjAiXSxbLTUwLCItIl0sWy01NCwie1wiaFwiOltcIl8zXCIsXCIzMjk5OTEzNjlcIixcIjY4NDQwMTg5NlwiLFwiMzI5OTkxMzY5XCIsXCIxMjM4NDExODczXCIsXCIxMTc0OTg5NTU5XCIsXCJfMlwiLFwiMjM0MzM1MjMxNVwiXSxcImRcIjpbXSxcImJcIjpbXCJfMVwiLFwiMjA2Njc4MTQxMFwiLFwiMzk2NTA1MDMzXCIsXCI3OTUzOTU3MDlcIixcIjExOTYyNzcwNTVcIixcIjkxODc4Mzk2OVwiLFwiMTY5ODUyMDA3OVwiLFwiNDIwNjA3OTA3M1wiLFwiMzYyNzIyOTcxM1wiLFwiMTE3NDk4OTU1OVwiXSxcInNcIjoxfSJdLFstNTksImRlbmllZCJdLFstNjMsIjAiXSxbLTY1LCItIl0sWy0xNSwiLSJdLFstMjgsImVuLVVTIl0sWy0zOSwiW1wiMjAwMzAxMDdcIiwwLFwiR2Vja29cIixcIk5ldHNjYXBlXCIsXCJNb3ppbGxhXCIsbnVsbCxudWxsLHRydWUsOCxmYWxzZSxudWxsLDAsZmFsc2UsdHJ1ZSxudWxsLDAsdHJ1ZSx0cnVlXSJdLFstNDQsIjAsMCwwLDUiXSxbLTcyLCJFeFU9Il0sWy0yLCIzMixlWW5XV1o5KytudnU5cE0zT21wZmVFTUNtRWFxUUVBaUYwcENqWVZwcm9XbFpkbEJVRTE0QUZsVlVSY1pXT2lpQmxzYUVJVWdPa1VCSklRbnFkdEVreU01bHl5bHVlOXUzOW5qIl0sWy0yOSwiLSJdLFstMzEsInRydWUiXSxbLTM1LCJbMTc0NjcyOTA3OTY0NCw3XSJdLFstNDEsIi0iXSxbLTQ2LCIwIl0sWy01MSwiLSJdLFstNTcsIldFMFpYVlphVEZSY1YwMFhXa3RjV0UxY2ZGVmNWRnhYVFJrUlVVMU5TVW9ERmhaYVhWY1hXbFpXVWxCY1ZWaE9GMVpMWGhaYVZsZEtYRmRORmx0Y1hBZ01XdzVhRkZzUENnc1VEUXdKWEJRQUNRa0tGQUJhQVZzUENWc0tXd0FPQVJaMlRYaE1UVlo3VlZaYVVoZFRTZ01BQXcwUEFSQVZXRTBaZUV0TFdFQVhYVndaRVZGTlRVbEtBeFlXVmxzWFgxQktVVXRXVzFaTlgxVldUbHhMRjFwV1ZCWlFGZzRPQ0EwS0FGaGNDQXNCV2c4TlgxOWNDd2xjRHdzTkR3c0JXbHNQV2c0QkYxTktBd2dEREE0UENnMFFGVmhOR1VzWkVWRk5UVWxLQXhZV1Zsc1hYMUJLVVV0V1cxWk5YMVZXVGx4TEYxcFdWQlpRRmc0T0NBMEtBQT09Il0sWy02NiwiZ2VvbG9jYXRpb24sY2h1YWZ1bGx2ZXJzaW9ubGlzdCxjcm9zc29yaWdpbmlzb2xhdGVkLHNjcmVlbndha2Vsb2NrLHB1YmxpY2tleWNyZWRlbnRpYWxzZ2V0LHNoYXJlZHN0b3JhZ2VzZWxlY3R1cmwsY2h1YWFyY2gsYmx1ZXRvb3RoLGNvbXB1dGVwcmVzc3VyZSxjaHByZWZlcnNyZWR1Y2VkdHJhbnNwYXJlbmN5LGRlZmVycmVkZmV0Y2gsdXNiLGNoc2F2ZWRhdGEscHVibGlja2V5Y3JlZGVudGlhbHNjcmVhdGUsc2hhcmVkc3RvcmFnZSxkZWZlcnJlZGZldGNobWluaW1hbCxydW5hZGF1Y3Rpb24sY2hkb3dubGluayxjaHVhZm9ybWZhY3RvcnMsb3RwY3JlZGVudGlhbHMscGF5bWVudCxjaHVhLGNodWFtb2RlbCxjaGVjdCxhdXRvcGxheSxjYW1lcmEscHJpdmF0ZXN0YXRldG9rZW5pc3N1YW5jZSxhY2NlbGVyb21ldGVyLGNodWFwbGF0Zm9ybXZlcnNpb24saWRsZWRldGVjdGlvbixwcml2YXRlYWdncmVnYXRpb24saW50ZXJlc3Rjb2hvcnQsY2h2aWV3cG9ydGhlaWdodCxjYXB0dXJlZHN1cmZhY2Vjb250cm9sLGxvY2FsZm9udHMsY2h1YXBsYXRmb3JtLG1pZGksY2h1YWZ1bGx2ZXJzaW9uLHhyc3BhdGlhbHRyYWNraW5nLGNsaXBib2FyZHJlYWQsZ2FtZXBhZCxkaXNwbGF5Y2FwdHVyZSxrZXlib2FyZG1hcCxqb2luYWRpbnRlcmVzdGdyb3VwLGNod2lkdGgsY2hwcmVmZXJzcmVkdWNlZG1vdGlvbixicm93c2luZ3RvcGljcyxlbmNyeXB0ZWRtZWRpYSxneXJvc2NvcGUsc2VyaWFsLGNocnR0LGNodWFtb2JpbGUsd2luZG93bWFuYWdlbWVudCx1bmxvYWQsY2hkcHIsY2hwcmVmZXJzY29sb3JzY2hlbWUsY2h1YXdvdzY0LGF0dHJpYnV0aW9ucmVwb3J0aW5nLGZ1bGxzY3JlZW4saWRlbnRpdHljcmVkZW50aWFsc2dldCxwcml2YXRlc3RhdGV0b2tlbnJlZGVtcHRpb24saGlkLGNodWFiaXRuZXNzLHN0b3JhZ2VhY2Nlc3Msc3luY3hocixjaGRldmljZW1lbW9yeSxjaHZpZXdwb3J0d2lkdGgscGljdHVyZWlucGljdHVyZSxtYWduZXRvbWV0ZXIsY2xpcGJvYXJkd3JpdGUsbWljcm9waG9uZSJdLFstNjcsIi0iXSxbLTY4LCJbMjg2NjAzNTQ0OSxmdW5jdGlvbigpe2Zvcih2YXIgYSxjLGI9W10sZT0wO2U8YXJndW1lbnRzLmxlbmd0aDtlKyspYltlXT1hcmd1bWVudHNbZV07cmV0dXJuXCJzY3JpcHRcIj09PWJbMF0udG9Mb3dlckNhc2UoKXx8LTEhPT10LmluZGV4T2YoYlswXS50b0xvd2VyQ2FzZSgpKT8oYT1CLmJpbmQoZG9jdW1lbnQpLmFwcGx5KHZvaWQgMCxiKSxjPWEuc2V0QXR0cmlidXRlLmJpbmQoYSksT2JqZWN0LmRlZmluZVByb3BlcnRpZXMoYSx7c3JjOnsiXSxbLTY5LCJNYWNJbnRlbHxHb29nbGUgSW5jLnw4fDEyfFdpbmRvd3N8MCJdLFsiYm5jaCIsMTM4XSxbLTQ3LCJBbWVyaWNhL0xvc19BbmdlbGVzLGVuLVVTLGxhdG4sZ3JlZ29yeSJdLFstNDgsIltcIi1cIixcIi1cIixcIi1cIl0iXSxbLTcwLCItIl0sWyJkZGIiLCIwLDMyLDAsMCwxLDAsMCwwLDAsMCwwLDAsMCwwLDEsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwxLDQsMCwxLDAsMCwwLDAsMCwwLDAsMCwwLDAsMSw2MiwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMSwwIl0sWyJjYiIsIjAsMCwwLDAsMCwwLDAsMSwwLDAsMCwwLDMsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMiwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDAsMCwwLDEiXV0%3D&amp;dep=0&amp;pre=0&amp;sdd=&amp;cri=jfECv66HI9&amp;pto=1512&amp;ver=65&amp;gac=462250241.1746729079&amp;mei=&amp;ap=&amp;fe=1&amp;duid=1.1746729079.FC7U58A4vZI5yYK0&amp;suid=1.1746729079.KGNJ6zRpGQhCdArz&amp;tuid=1.1746729079.hJyS90a8DamtJXqt&amp;fbc=1.1746729079520.75795469452619173&amp;gtm=WyJPbmVUcnVzdExvYWRlZCIsIk9wdGFub25Mb2FkZWQiLCJPbmVUcnVzdEdyb3Vwc1VwZGF0ZWQiXQ%3D%3D&amp;it=146%2C1052%2C109&amp;fbcl=-&amp;gacl=-&amp;gacsd=-&amp;rtic=-&amp;rtict=-&amp;bgc=-&amp;spa=1&amp;urid=0&amp;ab=jx.2.0%3B&amp;sck=-&amp;io=aGA2Oi17bmY2Og%3D%3D&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:417900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:4130,&#34;encoded_body_size&#34;:1487,&#34;decoded_body_size&#34;:4130,&#34;transfer_size&#34;:1787,&#34;download&#34;:{&#34;duration&#34;:900000,&#34;start&#34;:417000000},&#34;first_byte&#34;:{&#34;duration&#34;:120200000,&#34;start&#34;:296800000},&#34;connect&#34;:{&#34;duration&#34;:212800000,&#34;start&#34;:83900000},&#34;ssl&#34;:{&#34;duration&#34;:94800000,&#34;start&#34;:201900000},&#34;dns&#34;:{&#34;duration&#34;:44000000,&#34;start&#34;:39900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079905,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;44d2f8b7-e39d-4b53-8bc8-ac8e36bc1cdc&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.userway.org/widgetapp/2025-05-08-12-13-34/widget_app_base_1746706414131.js&#34;,&#34;status_code&#34;:200,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:176800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079863,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;68f463a2-c5bb-45e1-9dc8-78a88063a745&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:225000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:0,&#34;url&#34;:&#34;https://web.chtbl.com/track&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=b007d252-1f0a-4489-b1d6-f40469550ba5&amp;amp;batch_time=1746729080429", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079717,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c92b0ec4-305f-4cb6-963c-cbf5cc1d4ba0&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://c.contentsquare.net/pageview?ex=&amp;dt=134&amp;pvt=n&amp;cvars=%7B%221%22%3A%5B%22Page%20Name%22%2C%22%2Fen-us%2F%22%5D%2C%222%22%3A%5B%22Site%20ID%22%2C%22www.crowdstrike.com%22%5D%2C%224%22%3A%5B%22URL%22%2C%22www.crowdstrike.com%2Fen-us%2F%22%5D%2C%2211%22%3A%5B%22DB%20Industry%20Data%22%2C%22%25demandbaseDataElement1%25%22%5D%2C%2212%22%3A%5B%22DB%20Company%20Data%22%2C%22%25demandbaseDataElement2%25%22%5D%7D&amp;cvarp=%7B%221%22%3A%5B%22Page%20Name%22%2C%22%2Fen-us%2F%22%5D%2C%222%22%3A%5B%22Site%20ID%22%2C%22www.crowdstrike.com%22%5D%2C%224%22%3A%5B%22URL%22%2C%22www.crowdstrike.com%2Fen-us%2F%22%5D%2C%2211%22%3A%5B%22DB%20Industry%20Data%22%2C%22%25demandbaseDataElement1%25%22%5D%2C%2212%22%3A%5B%22DB%20Company%20Data%22%2C%22%25demandbaseDataElement2%25%22%5D%7D&amp;la=en-US&amp;uc=0&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;dr=&amp;dw=1280&amp;dh=7962&amp;ww=1280&amp;wh=720&amp;sw=1280&amp;sh=720&amp;uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;sn=1&amp;hd=1746729079&amp;v=15.85.1&amp;pid=29632&amp;pn=1&amp;r=089855&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:385400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:0,&#34;encoded_body_size&#34;:0,&#34;decoded_body_size&#34;:0,&#34;transfer_size&#34;:300,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:385300000},&#34;first_byte&#34;:{&#34;duration&#34;:83800000,&#34;start&#34;:301500000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729079726,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;2d5dc7eb-5e91-4f1f-b383-0ce3c9bfbda9&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://c.contentsquare.net/dvar?v=15.85.1&amp;pid=29632&amp;pn=1&amp;sn=1&amp;uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;dv=H4sIAAAAAAAAA6tWSi72TSxJzsjMS%2FdOrVSyUjLQszQyMDUxNTUyMTE2NDIxNI83NDcxMzeyNDC3NDQ1VaoFAC3Wr880AAAA&amp;ct=2&amp;r=397738&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:375700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:0,&#34;encoded_body_size&#34;:0,&#34;decoded_body_size&#34;:0,&#34;transfer_size&#34;:300,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:375500000},&#34;first_byte&#34;:{&#34;duration&#34;:83700000,&#34;start&#34;:291800000},&#34;connect&#34;:{&#34;duration&#34;:232200000,&#34;start&#34;:59300000},&#34;ssl&#34;:{&#34;duration&#34;:107700000,&#34;start&#34;:183800000},&#34;dns&#34;:{&#34;duration&#34;:38100000,&#34;start&#34;:21200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080060,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;0dd896f4-b72a-49f5-bce7-55a1c45ab5c9&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:50000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:204,&#34;url&#34;:&#34;https://px.ads.linkedin.com/wa/&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080106,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;afa69f60-87c7-4f06-a829-8dd69a31a725&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/u?mapType=mkto&amp;mapValue=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729080105&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=1&amp;a=crowdstrike.com&amp;rnd=598845&amp;cdn_o=a&amp;_biz_z=1746729080105&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:16500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080074,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;96d887bc-24c0-4e31-8191-def1c33f5499&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://td.doubleclick.net/td/rul/797629828?random=1746729080069&amp;cv=11&amp;fst=1746729080069&amp;fmt=3&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gcl_ctr=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;label=hozuCPn52LoYEIS7q_wC&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;gtm_ee=1&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;capi=1&amp;data=event%3Dconversion&amp;ct_cookie_present=0&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:75500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:3989,&#34;encoded_body_size&#34;:1199,&#34;decoded_body_size&#34;:3989,&#34;transfer_size&#34;:1499,&#34;download&#34;:{&#34;duration&#34;:1200000,&#34;start&#34;:74300000},&#34;first_byte&#34;:{&#34;duration&#34;:73100000,&#34;start&#34;:1200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080076,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6cb8b9b6-786e-44bc-bfba-5c4878f713b9&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://obs.fishrobotflower.com/tracker/tc_imp.gif?e=37dfbd8ee84e00126ee8c037e347829d9225c24f567d43d6da1908be6245cad7bd70a976750ef80ed89373bfe70e9c20c1e53e8d5a138c6c2717071a10acf9f29f671a82d589562b6d1ef878215780338a33c103300376c5065359600d5ccfbf394f77be26bb25cb43e2913df05365ac0b297d1bdc57e446f492d7da3dbb280ef67ccaa8073f8d0e3717224793845731f360b3f493a0180dec1edae97dfa2bc8169b1adc597cff3200e714561c4b92177af998ffe4198b6dec06c213f85e162ae7d133722b325f817c99ec59b058609fc6e359143e3dd385293e88864c06513c157a77bb9e70392652b48d1c2ad7f4ec3ee3b8192d4079b4a7a6928677a0dadf5ae8489e5d2e019cbecbf7af2b95dfe57594351ccdeb8b795904fd7367a095166dbdd0aa73d902fd12b0801d91499d9978953b67919d27dc6796d3bbe193fdbd4c38fc28b0bdfd354371fe8f719aa61af7010642dd4245c2979684c0ec8825ce3bc68ebe817fa27d763bc97f27db838003679ea875ac41bfc003d667a87e5346588c6eee5f8b89ee7ef8df78f02da7c19cdb27d60a0a2dea5d540b28073ed7ab67b355c4cb9764fdde25b3d02dd0c2b1be7c85b7f409746c89a91004d7b41d5107f7fae939fa2890f802fdfc3ce52dd36dcc068a50aefc2ecd0823e2ac90633f3d6402340a5b13ace68e65d205be5752517faac71386cd464887eccdb0ae6ddcece3f880ffadf84b4835fe4586a85e6b844a601fd16a59746ef07aa6771ec1f9dd71c83cc52fe0a82cc8a40e99721322210760129cc0e4db207a559a184689d2d9bbb0c24f049eb0be7f8ad2757adbce0f3a68fd8dc45020b23b98599e1cff22981ad8d40efbb337fe868bd49e5b7feee3337f9f71a68aacb6119d6d280b02956a5f4db6a25b93f112ca9887d552cd839b2e083745d321b951ef0ce6936384335d20463e7b78e3538e863a4bcf1cc7c2f03da8cb0be60d76e678072c8df4c3491fb725e78402db1930074eba3c2e8bafd061d52c49e4ff15ebbd9d88e11aa90924605fe06b3716d9c9584eb02ee6e0c02a8bd5eb0c547ca5633e507e867459e6c8e66c40cd260bbbd6c353d5c6f88f0f30b05bd61b781488173175d4cc1615da92b2aabde142a563f930b77096d65d9afff22a2bfdd215808cb4b1c4f6cadf741fddf42f54872dd257c81cf3ee28a8cad1c731efc4c7f2fb5a69cba2b5af51cb5594404921413d88823335dc60f64f8ac965ef0522e2714c9766e69c7575b97521525796f3d12f87c74c603d1e6ac2d583ffce748a07e2ef4b7ee4de495b5018376e2de975e2a73802b3&amp;cri=jfECv66HI9&amp;ts=429&amp;cb=1746729080075&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:100400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080069,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;da2e5204-fdf9-4cf5-abb2-f40d3b6d02f5&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://bat.bing.com/bat.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:149400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080208,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;699dbe1a-40d2-4e4e-91a5-dcf440671332&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/u?mapType=ecid&amp;mapValue=06D71E9261F941560A495CD6%40AdobeOrg_28180928486187528650392535887012291588&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729080106&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=2&amp;a=crowdstrike.com&amp;rnd=228061&amp;cdn_o=a&amp;_biz_z=1746729080208&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:47100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080175,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8d97b5a3-00f0-4644-9414-f887d6e73e94&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://c.contentsquare.net/dvar?v=15.85.1&amp;pid=29632&amp;pn=1&amp;sn=1&amp;uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;dv=H4sIAAAAAAAAA3WSX0%2BDMBTFv0pDfBwG5uZ0b4VtkWR%2FklUT4wupcMFGaElbtqHxu9tmOmDTR%2Fo79%2FScWz4dHMT4MX4QJVQ0B7SInudkkLGDu1favRl5rt%2F7GjpTpy%2Bezg8VSAY8AYT7k%2F%2BjYQ85AxNjFd%2BSeCZKyri547WWOoMiLSm%2FTkTZKkJRcy0bFJENCkUKRvtEWow5r2mBtrADXhs4nkyGnued%2BBpHITEz45Hv%2B57TP0czUIlklWbCZthkGUtAIZGhJd03IFWrJ1FoFHfGpHd25rCE3IQhIHfWqFuirChv0JqWtkBgyqLFse1FU1uQMw0pIprqngvTlpJEaK1SWkAnipUahiX7EJyegd%2FF4ZcWzMuqEA1YZq51pv4l2VKe2zEPuei%2BnYx4Wqtj0KBWjINSf1T%2BeZGTydV4ZWyufG%2FV1eTHva2F1G8gOcKl%2BUmSTv5IoaCgyXvBlFmJkWa0UOB8fQN6Gti9yQIAAA%3D%3D&amp;ct=2&amp;r=414786&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:87700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:0,&#34;encoded_body_size&#34;:0,&#34;decoded_body_size&#34;:0,&#34;transfer_size&#34;:300,&#34;download&#34;:{&#34;duration&#34;:300000,&#34;start&#34;:87400000},&#34;first_byte&#34;:{&#34;duration&#34;:87200000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080086,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1be8a31c-e5b5-4c3f-b071-8a2fc28ea55a&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:191000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://api.userway.org/api/v1/tunings/dyvvHf6oG0&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080220,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;239bbd0a-6400-47bc-b727-447c92edb399&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://bat.bing.com/p/action/163002607.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:59400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080278,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;bf69cd19-c29a-42fc-a3dd-9d788941fa2a&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:27000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://cdn.userway.org/widgetapp/2025-05-08-12-13-34/locales/en-US.json&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.google.com/pagead/1p-conversion/797629828/?random=160478099&amp;amp;cv=11&amp;amp;fst=1746729080069&amp;amp;bg=ffffff&amp;amp;guid=ON&amp;amp;async=1&amp;amp;gcl_ctr=1&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;gcd=13l3l3l3l1l1&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;u_w=1280&amp;amp;u_h=720&amp;amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;label=hozuCPn52LoYEIS7q_wC&amp;amp;hn=www.googleadservices.com&amp;amp;frm=0&amp;amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;did=dYWJhMj&amp;amp;gdid=dYWJhMj&amp;amp;gtm_ee=1&amp;amp;npa=0&amp;amp;pscdl=noapi&amp;amp;auid=1475816545.1746729079&amp;amp;uaa=x64&amp;amp;uab=64&amp;amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;amp;uamb=0&amp;amp;uam=&amp;amp;uap=Windows&amp;amp;uapv=10.0&amp;amp;uaw=0&amp;amp;fledge=1&amp;amp;capi=1&amp;amp;data=event%3Dconversion&amp;amp;fmt=3&amp;amp;ct_cookie_present=false&amp;amp;crd=CPLOsQIIobixAgixwbECCLDBsQIIscOxAgiKxbECCMLJsQIItMaxAgiQybECCNPFsQII68yxAgjPzrECCP7OsQII1c-xAkoVdHJpZ2dlciwgZXZlbnQtc291cmNlWgMKAQFiBAoCAgM&amp;amp;pscrd=IhMI34Lk8sCUjQMVqKOOCB2_0BUQMgwIA2IICAAQABgAIAAyDAgEYggIABAAGAAgADIMCAdiCAgAEAAYACAAMgwICGIICAAQABgAIAAyDAgJYggIABAAGAAgADIMCApiCAgAEAAYACAAMgwIAmIICAAQABgAIAAyDAgLYggIABAAGAAgADIMCBViCAgAEAAYACAAMgwIH2IICAAQABgAIAAyDAgTYggIABAAGAAgADIMCBJiCAgAEAAYACAAOhxodHRwczovL3d3dy5jcm93ZHN0cmlrZS5jb20vQlZDaEVJOEt6eHdBWVFvTE9vdXAtQXU5bjlBUklyQVBSSTZXMURUUl94MUZMVUdqZUNFdGtHbWhGY0RkQmdnRHdlMVQ4SUNUYVpZRTdHN095bFlqY0ZBUQ&amp;amp;is_vtc=1&amp;amp;cid=CAQSKQDZpuyzUTR6hnjEyRsk0u0adWtaKTCvPq0x3tHMxzJNy2AH9tqYktcM&amp;amp;random=565709458", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core?d=1&amp;amp;embedId=9d4udx6ceimp&amp;amp;eId=9d4udx6ceimp&amp;amp;region=US&amp;amp;forceShow=false&amp;amp;skipCampaigns=false&amp;amp;sessionId=bb076d85-176d-40bf-8edd-e4d09c550682&amp;amp;sessionStarted=1746729080.521&amp;amp;campaignRefreshToken=f7d1fcdd-1c5b-41be-9318-104dbcaa8709&amp;amp;hideController=false&amp;amp;pageLoadStartTime=1746729078294&amp;amp;mode=CHAT&amp;amp;driftEnableLog=false&amp;amp;secureIframe=false&amp;amp;u=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/chat?d=1&amp;amp;region=US&amp;amp;driftEnableLog=false&amp;amp;pageLoadStartTime=1746729078294", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://edge.adobedc.net/ee/or2/v1/interact?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=67afc763-a8cd-4361-9eb8-ad6d68e97d82", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;events&#34;:[{&#34;xdm&#34;:{&#34;eventType&#34;:&#34;web.webinteraction.linkClicks&#34;,&#34;eventMergeId&#34;:&#34;75cfb0ab-4767-47ed-91e9-50bfe24f7ea1&#34;,&#34;web&#34;:{&#34;webPageDetails&#34;:{&#34;URL&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;},&#34;webReferrer&#34;:{&#34;URL&#34;:&#34;&#34;},&#34;webInteraction&#34;:{&#34;URL&#34;:&#34;/en-us/&#34;,&#34;name&#34;:&#34;DFP Event&#34;,&#34;type&#34;:&#34;other&#34;,&#34;linkClicks&#34;:{&#34;value&#34;:1}}},&#34;device&#34;:{&#34;screenHeight&#34;:720,&#34;screenWidth&#34;:1280,&#34;screenOrientation&#34;:&#34;landscape&#34;},&#34;environment&#34;:{&#34;type&#34;:&#34;browser&#34;,&#34;browserDetails&#34;:{&#34;viewportWidth&#34;:1280,&#34;viewportHeight&#34;:720}},&#34;placeContext&#34;:{&#34;localTimezoneOffset&#34;:420,&#34;localTime&#34;:&#34;2025-05-08T11:31:20.532-07:00&#34;},&#34;timestamp&#34;:&#34;2025-05-08T18:31:20.532Z&#34;,&#34;implementationDetails&#34;:{&#34;name&#34;:&#34;https://ns.adobe.com/experience/alloy/reactor&#34;,&#34;version&#34;:&#34;2.26.0&#43;2.29.0&#34;,&#34;environment&#34;:&#34;browser&#34;},&#34;_experience&#34;:{&#34;analytics&#34;:{&#34;event101to200&#34;:{&#34;event102&#34;:{&#34;id&#34;:&#34;DFP Event&#34;,&#34;value&#34;:1}},&#34;customDimensions&#34;:{&#34;eVars&#34;:{&#34;eVar1&#34;:&#34;/en-us/&#34;,&#34;eVar2&#34;:&#34;www.crowdstrike.com&#34;,&#34;eVar3&#34;:&#34;&#34;,&#34;eVar4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;eVar6&#34;:&#34;dir&#34;,&#34;eVar7&#34;:&#34;&#34;,&#34;eVar10&#34;:&#34;35a249abd9e94512be1326ac06e4ddd5&#34;,&#34;eVar11&#34;:&#34;id:281-OBQ-266&amp;token:_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&#34;,&#34;eVar13&#34;:&#34;fb.1.1746729079520.75795469452619173&#34;,&#34;eVar14&#34;:&#34;GA1.1.462250241.1746729079&#34;,&#34;eVar15&#34;:&#34;28180928486187528650392535887012291588&#34;,&#34;eVar16&#34;:&#34;bg41,c0001,c0002,c0003,c0004&#34;,&#34;eVar17&#34;:&#34;&#34;,&#34;eVar18&#34;:&#34;&#34;,&#34;eVar19&#34;:&#34;CrowdStrike-AEM|production|2025-05-05T15:24:36Z|Analytics: Device Finger Print | Window Load | Max Frequency -1 | Event-102&#34;,&#34;eVar20&#34;:&#34;/en-us/&#34;,&#34;eVar26&#34;:&#34;in&#34;,&#34;eVar111&#34;:&#34;crowdstrike:homepage&#34;,&#34;eVar112&#34;:&#34;crowdstrike:en-us&#34;,&#34;eVar113&#34;:&#34;&#34;,&#34;eVar114&#34;:&#34;&#34;,&#34;eVar115&#34;:&#34;crowdstrike:en-us&#34;},&#34;lists&#34;:{&#34;list3&#34;:{&#34;list&#34;:[{&#34;value&#34;:&#34;BrowserFlag : 9&#34;},{&#34;value&#34;:&#34;HasBattery : 1&#34;},{&#34;value&#34;:&#34;BatteryLevel : 99&#34;},{&#34;value&#34;:&#34;ScrollbarWidth : 0&#34;},{&#34;value&#34;:&#34;ChannelCountMode : explicit&#34;},{&#34;value&#34;:&#34;MaxChannelCount : 2&#34;},{&#34;value&#34;:&#34;Web3.isConnected : null&#34;},{&#34;value&#34;:&#34;Web3.networkVersion : null&#34;},{&#34;value&#34;:&#34;Web3.chainId : null&#34;},{&#34;value&#34;:&#34;ColorDepth : 24&#34;},{&#34;value&#34;:&#34;PixelRatio : 1&#34;},{&#34;value&#34;:&#34;PixelDepth : 24&#34;},{&#34;value&#34;:&#34;WindowScreenWidth : 1280&#34;},{&#34;value&#34;:&#34;WindowScreenHeight : 720&#34;},{&#34;value&#34;:&#34;WindowScreenAvailableWidth : 1280&#34;},{&#34;value&#34;:&#34;WindowScreenAvailableHeight : 720&#34;},{&#34;value&#34;:&#34;JsHeapSizeLimit : 3760000000&#34;},{&#34;value&#34;:&#34;OsTheme : light&#34;},{&#34;value&#34;:&#34;Browser : Chromium&gt;HeadlessChrome&gt;Not.A/Brand&#34;},{&#34;value&#34;:&#34;InstalledFontHash : 504800432572425&#34;},{&#34;value&#34;:&#34;InstalledFontNames : AcademyEngravedLetPlain,American Typewriter,Apple Chancery,Apple SD Gothic Neo,Apple Symbols,AppleGothic,AppleMyungjo,Arial Black,Arial Narrow,Arial Rounded MT Bold,Avenir Next Condensed,Avenir Next,Avenir,Ayuthaya,Bangla MN,Bangla Sangam MN,Baskerville,Big Caslon,Bodoni Ornaments,Bradley Hand,Brush Script MT,Chalkboard SE,Chalkboard,Chalkduster,Charter,Cochin,Comic Sans MS,Consolas (supports Cyrillic),Copperplate,Corbel (supports Cyrillic),DIN Alternate,DIN Condensed,Devanagari Sangam MN,Didot,Euphemia UCAS,Futura,Galvji,Geneva,Georgia,Gill Sans,Grantha Sangam MN,Gujarati Sangam MN,Gurmukhi MN,Gurmukhi Sangam MN,Heiti SC,Heiti SC/TC,Heiti TC,Helvetica Neue,Herculanum,Hiragino Maru Gothic ProN,Hiragino Mincho ProN,Hiragino Sans GB,Hiragino Sans,Hoefler Text,Impact,InaiMathi,Kannada MN,Kannada Sangam MN,Kefa,Khmer MN,Khmer Sangam MN,Kohinoor Bangla,Kohinoor Devanagari,Kohinoor Telugu,Krungthep,Lao MN,Lao Sangam MN,LiHei Pro,LiSong Pro,Lucida Grande,Luminari,Malayalam MN,Malayalam Sangam MN,Marker Felt,Menlo,Microsoft Sans Serif,Myanmar MN,Myanmar Sangam MN,Nanum Brush Script,Nanum Gothic,Nanum Myeongjo,Nanum Pen Script,Noteworthy,Optima,Oriya MN,Oriya Sangam MN,PCMyungjo,PT Mono,PT Sans Caption,PT Sans Narrow,PT Sans,PT Serif Caption,PT Serif,Palatino Linotype (supports Cyrillic),Palatino,Papyrus,Party LET,Phosphate,PingFang HK,PingFang SC,PingFang SC/TC,PingFang TC,Plantagenet Cherokee,Rockwell,STFangsong,STHeiti,STIXGeneral,STIXVariants,STKaiti,STSong,STXihei,Sathu,Savoye LET,SignPainter,Silom,Sinhala MN,Sinhala Sangam MN,Skia,Snell Roundhand,Songti SC,Songti SC/TC,Songti TC,Sukhumvit Set,Tahoma,Tamil MN,Tamil Sangam MN,Telugu MN,Telugu Sangam MN,Thonburi,Trattatello,Trebuchet MS,Verdana,Webdings,Wingdings,Xingkai SC,Zapfino,academy engraved let,cursive,fantasy&#34;},{&#34;value&#34;:&#34;Languages : en-US&#34;},{&#34;value&#34;:&#34;Network.downlink : 10&#34;},{&#34;value&#34;:&#34;Network.effectiveType : 4g&#34;},{&#34;value&#34;:&#34;Network.isThrottled : false&#34;},{&#34;value&#34;:&#34;Network.rtt : null&#34;},{&#34;value&#34;:&#34;Network.downlinkMax : null&#34;},{&#34;value&#34;:&#34;Network.type : null&#34;},{&#34;value&#34;:&#34;DeviceMemory : 8&#34;},{&#34;value&#34;:&#34;Platform : Windows&#34;},{&#34;value&#34;:&#34;TimeZone : 420&#34;},{&#34;value&#34;:&#34;TimeZoneLabel : America/Los_Angeles&#34;},{&#34;value&#34;:&#34;VideoCardVendor : Google Inc. (Apple)&#34;},{&#34;value&#34;:&#34;VideoCardRenderer : ANGLE (Apple, Apple M4 Pro, OpenGL 4.1)&#34;},{&#34;value&#34;:&#34;VideoCardVersion : WebGL 1.0 (OpenGL ES 2.0 Chromium)&#34;},{&#34;value&#34;:&#34;ShadingLanguageVersion : WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)&#34;},{&#34;value&#34;:&#34;Architecture : x64&#34;},{&#34;value&#34;:&#34;BrowserFullVersion : 136.0.7103.25&#34;},{&#34;value&#34;:&#34;PermissionsFlag : null&#34;},{&#34;value&#34;:&#34;Plugins&#34;},{&#34;value&#34;:&#34;ScreenFlag : 0&#34;}]}},&#34;props&#34;:{&#34;prop1&#34;:&#34;/en-us/&#34;,&#34;prop2&#34;:&#34;www.crowdstrike.com&#34;,&#34;prop4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;prop6&#34;:&#34;&#34;,&#34;prop10&#34;:&#34;crowdstrike:homepage&#34;,&#34;prop11&#34;:&#34;crowdstrike:en-us&#34;,&#34;prop12&#34;:&#34;&#34;,&#34;prop13&#34;:&#34;&#34;,&#34;prop14&#34;:&#34;crowdstrike:en-us&#34;}}}}}}],&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true,&#34;entries&#34;:[{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_identity&#34;,&#34;value&#34;:&#34;CiYyODE4MDkyODQ4NjE4NzUyODY1MDM5MjUzNTg4NzAxMjI5MTU4OFISCJ2jqonrMhABGAEqA09SMjAA8AGdo6qJ6zI=&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_cluster&#34;,&#34;value&#34;:&#34;or2&#34;}]}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.zi-scripts.com/zi-tag.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/rtp/api/v1_1/visitor?sid=crowdstrike-1746729079515-66d02f45&amp;amp;aid=crowdstrike&amp;amp;1746729080540", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/ga/sgm?sid=crowdstrike-1746729079515-66d02f45&amp;amp;1746729080540", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://px.ads.linkedin.com/wa/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;pids&#34;:[64444],&#34;scriptVersion&#34;:213,&#34;time&#34;:1746729080586,&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;url&#34;:&#34;https://crowdstrike.com/en-us/&#34;,&#34;pageTitle&#34;:&#34;CrowdStrike: We Stop Breaches with AI-native Cybersecurity&#34;,&#34;websiteSignalRequestId&#34;:&#34;41aa6e30-c3ae-a9cf-9bec-305037e2872c&#34;,&#34;isTranslated&#34;:false,&#34;liFatId&#34;:&#34;&#34;,&#34;liGiant&#34;:&#34;&#34;,&#34;misc&#34;:{&#34;psbState&#34;:-4},&#34;isLinkedInApp&#34;:false,&#34;hem&#34;:null,&#34;signalType&#34;:&#34;CLICK&#34;,&#34;href&#34;:&#34;&#34;,&#34;domAttributes&#34;:{&#34;elementSemanticType&#34;:&#34;ICON&#34;,&#34;elementValue&#34;:null,&#34;elementType&#34;:null,&#34;tagName&#34;:&#34;BUTTON&#34;,&#34;backgroundImageSrc&#34;:&#34;https://cdn.cookielaw.org/logos/static/ot_close.svg&#34;,&#34;imageSrc&#34;:null,&#34;imageAlt&#34;:null,&#34;innerText&#34;:&#34;&#34;,&#34;elementTitle&#34;:&#34;Close&#34;,&#34;cursor&#34;:&#34;pointer&#34;},&#34;innerElements&#34;:null,&#34;elementCrumbsTree&#34;:[{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:15,&#34;id&#34;:&#34;onetrust-consent-sdk&#34;},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:1,&#34;id&#34;:&#34;onetrust-banner-sdk&#34;,&#34;classes&#34;:[&#34;bottom&#34;,&#34;otFlat&#34;,&#34;vertical-align-content&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:1,&#34;id&#34;:&#34;onetrust-close-btn-container&#34;},{&#34;tagName&#34;:&#34;button&#34;,&#34;nthChild&#34;:0,&#34;classes&#34;:[&#34;banner-close-button&#34;,&#34;onetrust-close-btn-handler&#34;,&#34;onetrust-close-btn-ui&#34;,&#34;ot-close-icon&#34;]}],&#34;isFilteredByClient&#34;:false}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.facebook.com/tr/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;id&#34;

992980065451679
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;ev&#34;

SubscribedButtonClick
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;dl&#34;

https://www.crowdstrike.com/en-us/
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;rl&#34;


------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;if&#34;

false
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;ts&#34;

1746729080589
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;cd[buttonFeatures]&#34;

{&#34;classList&#34;:&#34;onetrust-close-btn-handler onetrust-close-btn-ui banner-close-button ot-close-icon&#34;,&#34;destination&#34;:&#34;&#34;,&#34;id&#34;:&#34;&#34;,&#34;imageUrl&#34;:&#34;url(\&#34;https://cdn.cookielaw.org/logos/static/ot_close.svg\&#34;)&#34;,&#34;innerText&#34;:&#34;&#34;,&#34;numChildButtons&#34;:0,&#34;tag&#34;:&#34;button&#34;,&#34;type&#34;:null,&#34;name&#34;:&#34;&#34;,&#34;value&#34;:&#34;&#34;}
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;cd[buttonText]&#34;


------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;cd[formFeatures]&#34;

[]
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;cd[pageFeatures]&#34;

{&#34;title&#34;:&#34;CrowdStrike: We Stop Breaches with AI-native Cybersecurity&#34;}
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;sw&#34;

1280
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;sh&#34;

720
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;v&#34;

2.9.201
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;r&#34;

stable
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;ec&#34;

1
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;o&#34;

4126
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;fbp&#34;

fb.1.1746729079520.75795469452619173
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;ler&#34;

empty
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;it&#34;

1746729079168
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;coo&#34;

false
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;es&#34;

automatic
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;tm&#34;

3
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;exp&#34;

k2
------WebKitFormBoundaryNsANUl396lwtNArQ
Content-Disposition: form-data; name=&#34;rqm&#34;

SB
------WebKitFormBoundaryNsANUl396lwtNArQ--
`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://privacyportal.onetrust.com/request/v1/consentreceipts", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;requestInformation&#34;:&#34;eyJhbGciOiJSUzUxMiJ9.eyJvdEp3dFZlcnNpb24iOjEsInByb2Nlc3NJZCI6IjgzNTFlYTFlLWJjMDItNDM0ZC04N2M5LTk2MTE1MTcwZGUyNSIsInByb2Nlc3NWZXJzaW9uIjo1MiwiaWF0IjoiMjAyMi0xMi0yMVQxNzoyOToyNS44OSIsIm1vYyI6IkNPT0tJRSIsInBvbGljeV91cmkiOiJjcm93ZHN0cmlrZS5jb20iLCJzdWIiOiJDb29raWUgVW5pcXVlIElkIiwiaXNzIjpudWxsLCJ0ZW5hbnRJZCI6ImMxMDlkYWU5LTQ2ZjMtNGU5MS1hNTllLTc4NDRlZjY0NTEwNyIsImRlc2NyaXB0aW9uIjoiVGhpcyBjb2xsZWN0aW9uIHBvaW50IGNhcHR1cmVzIHRoZSBjdXJyZW50IHNpdGUgdmlzaXRvciBjb25zZW50IHByZWZlcmVuY2VzIGZvciB0aGUgZG9tYWluOiBjcm93ZHN0cmlrZS5jb20tZW4gc3BlY2lmaWVkLiIsImNvbnNlbnRUeXBlIjoiQ09PS0lFQkFOTkVSIiwiYWxsb3dOb3RHaXZlbkNvbnNlbnRzIjpmYWxzZSwiZG91YmxlT3B0SW4iOmZhbHNlLCJwdXJwb3NlcyI6W3siaWQiOiI4ZDU0ZDBhOC04NTBiLTRlNjctYTc4ZC1mYmFhN2M3M2JlMGYiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiIzZjI0YjgzOS05YTI5LTQ3M2EtYjExNC1jYzQzY2E4Njg1ZTAiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI1N2MwZjNlNy01NGQ1LTRiOTItYjNiMC1jYWRiMWNhNjBkZDAiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI0NTYzYjA5NC0yNWZjLTQ2ZDEtYmVmMS1kMzcwZTBmMGJlMjEiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI4YWNkMjg4NS1jN2FjLTRhMTgtYmE1YS03YTBiOTQ2NmU4OGYiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI4MjA2MTdjZi02ZTgzLTQ5NmUtODcxZC01ZTFkY2M4N2YwZGMiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI0OTNhYjkwYy1mNzQ2LTQ4N2YtYjhkOC1kZThlMTQ2NmVjOTEiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI3YTVjOTVjMi1hZWJjLTQzNzQtYmI4ZC01MWE3ZmEwYzg4YWIiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI0ZjFhMmU3My1hOGQ5LTQ2NjgtYmJmOS03N2ZjNzdkNDJkZGUiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI1MDZlMGEzOS0xZmY4LTQ4MWItOTA1Zi02ZTZjY2E2NTVlYjYiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI3MjJhZmM0YS1hYTMwLTQzZWQtOWY5Yy05OTI4ZjkwNTkwNzEiLCJ2ZXJzaW9uIjozLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI5ZDkyYzc2NC04ZjM0LTRkMDMtYjYxNC1hYTFiOGFmYThiMmYiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJmMzBkOTlmNS00ZmFkLTQzMTctYjI0Ni1hN2ZmMmMyMmE0NTIiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI2YjRmMmM2OC01MGMxLTQ3YjUtOTRlMi00ZjI4MTM0MGI4MmUiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJkMGRkYmUzOS0wZGEyLTQ1Y2UtOWViNS0yZGQ3MTU4YWQ3YjAiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI1MjI4NmM5ZS0yYTgzLTRjNGEtOTRkYS05ODZjMmZjYTBiMjMiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiIxMDE2NmIxYi1mMzdhLTQ5YTMtYTlmNC1kOGFkNzA2NDg5YmMiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJjOTdkNDQ5OS04YzY5LTQ4MDgtOTAwNS1kODQ2YWU2NDRkMWUiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI4OGMyODZlOS1kMDBiLTRkMmEtODE3Yi1iZmU4OTI3ZDc1ZDciLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI5OGM4ZWRkOS00MmYwLTQxODQtYjI1NS0yODY0Y2ZiYTlhMTciLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJmNDM2YjU1NS1kOWE1LTQyZTctOTMwYy0zN2Q0MDc1MTljYjgiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJjYWFlNjlhMC1lNGMzLTRkYTktYjgxYi1kZjIxZmMwMjFhYTIiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiIyNThiZWU2Mi04NjE3LTRmNzgtOWM3OS04Yjg5MzMxN2ZlYjUiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJjNGU4OGU0Yy00NzE5LTQ2ZGMtOGUwOS04Y2M5NGFlOGQ4NTAiLCJ2ZXJzaW9uIjoyLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJkZGFkMzViYy1hZjllLTQyZDItYWVjOC1kNDM5N2UyZTg5NzIiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiIzMTFhODczNi0xZDViLTRkYzAtOTYyZC1jYTJlODVlYTA4MWIiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJlNjQ0MmQ3ZC03NmI3LTQxMzctOGNlNS03MzVhYjlmODgxMmMiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI0ZWJjMTFkZS1lMDY5LTQyZmUtOGVkZS01YzUxMjNhOTkwYzkiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJlZDJiMzI5MC1jNTcyLTRjZDctOGU0MC1kNDBhOTM5Zjk2ZDciLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJiOGIyN2JmZS1mNTIyLTQ2ODctODZkOS03NmUzMjdmN2RkNWQiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiIwZmE1M2ZkNy1mNTA2LTQ5OGQtYmY4MS1jOTZkZDFhMTU2OTQiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJmZTI3M2YwYi1mNGY0LTRhZDktYWQ2MS02NDI4OTE0YTQ5YTkiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI1MjA4YjM5NC0zNWUzLTQ5YmYtYjAyYS1mYjRjM2MwZWVmOGIiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiIwM2VmN2I2NC0xMjIyLTQ0MTQtODgxNy1iY2U3MzU4Nzk3ZTQiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI4YWM4MTkxYi1hMjYzLTRjN2UtYmY1ZC1jY2UyMDkxY2Y5N2QiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiIxYTc5MTk4Mi1hMGRkLTRlODYtOWRiYy0yYjg1NGQzYTI5NjYiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiIxODc1MDU1NC02YjBiLTQxZmUtOGM3Zi0wYzdmZDFiMjY4ZjQiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJhYmYyNjgwOC1hNDM4LTQ1OGItOGQ3MC0wYmQ2MTFiMzJhZWQiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI2NzgzMmJlYS0zNDAzLTRhZTAtYmNiNC1hZDk2MjljMmI2MGUiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJiYTViNzM2MC1mMWZiLTQyMjQtOGZkYi1iZGY5ZjQ4ZDFiMzkiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJiMzQ4ZmY2OS1mMjMzLTRkNmYtOGU4My1lNzgxMmZmMjU3NWEiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI4YjIzYWRmYS1lOWVlLTQ5N2YtOGI0Yi01MzY4Y2ZkZmVkZjYiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJmOTcyMzEwMC1hODY3LTRmMWItOWM5Yi01MGE3MDdkM2NlZTciLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJkYjMyZDg4ZC05MDlkLTQ1ZDMtOTVhMC04Njc4NWNjZTQ1MTUiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI0MmJmN2Y4NS1iZWFlLTQ3YjYtOGIxYi03OGI4ZTM1YThkY2IiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJkN2FiMjM2Mi1hOTczLTRiZTEtYjE2NC05ZGFmZmUwYjg5OGEiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiJiMDdkMTlhMC1iNmJmLTQ2MTMtYmRlNi1kZWZkMzNhNTE2YTEiLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9LHsiaWQiOiI2NmM2ZmRiZi0wYWE2LTQ0ZjUtODE3YS03N2JlMmNmM2U4MjciLCJ2ZXJzaW9uIjoxLCJwYXJlbnRJZCI6bnVsbCwidG9waWNzIjpbXSwiY3VzdG9tUHJlZmVyZW5jZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9XSwibm90aWNlcyI6W10sImRzRGF0YUVsZW1lbnRzIjpbIlVzZXJBZ2VudCIsIkNvdW50cnkiLCJJbnRlcmFjdGlvblR5cGUiLCJDb25zZW50TW9kZWwiXSwiYXV0aGVudGljYXRpb25SZXF1aXJlZCI6ZmFsc2UsInJlY29uZmlybUFjdGl2ZVB1cnBvc2UiOmZhbHNlLCJvdmVycmlkZUFjdGl2ZVB1cnBvc2UiOnRydWUsImR5bmFtaWNDb2xsZWN0aW9uUG9pbnQiOmZhbHNlLCJhZGRpdGlvbmFsSWRlbnRpZmllcnMiOltdLCJtdWx0aXBsZUlkZW50aWZpZXJUeXBlcyI6ZmFsc2UsImVuYWJsZVBhcmVudFByaW1hcnlJZGVudGlmaWVycyI6ZmFsc2UsInBhcmVudFByaW1hcnlJZGVudGlmaWVyc1R5cGUiOm51bGwsImFkZGl0aW9uYWxQYXJlbnRJZGVudGlmaWVyVHlwZXMiOltdLCJlbmFibGVHZW9sb2NhdGlvbiI6ZmFsc2V9.Va2IQmAVDMj9434sgEpksZIzoZ_YThsRVfnH0cg0ZU5O2KiH5tvvoa1RfUnYWFo_5rrb7X-eqA2JdE_U48Zo3BzD9Ci2LgIsSab6qWm1q-_DMgp38BfIvBslGN4hBFB35LvmInmNvrOHqasTNc98Mvu4ug-CR0rF_NSKGEmS_byH-Ngjrt2peo3wUM2dZ3ffc543gZidYvVxzbTlbaz5bhL61Noz3DGo9CKvfoAkxlB5iVzxokLIEgjr8aPYqDnk3H6O1poP-tcoC_zwoV8x9WV_HQdX7I7_gf2vlfOM04jVZsp8bGegauTshPbT184FfF032UqX9vDqzZWmE3aJNlpa162fLT1Pf1a70_NlQKnTRtVsLRZIpsOIxnBSfK_VpECQFiQM5wgawRezX7-lp7YLlj1O-QFxpscMVX7gjILi6KoBfYGRfpeovWxM-0bzNYbz3tohCkS9t3vZuH9wepQ18LtsOk0KuY7LPWpUOFFUZSJxBpDjajcsQ7uiS25hp7B63DWtMM0k9A0Lo3zjxEeG17bGT3dCfXUMd6YoAnStP5LQsQb13_8hw5LAT1FRVlm0oqjkGN06PNtMwxAohvKDAfyXL4N0-bEA67lf7i8KEWgxEG9I4mH-RZIVg5Tpn-hKW1C9SQ5eyJI3OpKtEl0_SElw9K0mjtbp-oRmRcA&#34;,&#34;identifier&#34;:&#34;6fab6eba-becf-476a-a133-42d6d5e77134&#34;,&#34;identifierType&#34;:&#34;Cookie Unique Id&#34;,&#34;customPayload&#34;:{&#34;Interaction&#34;:1,&#34;AddDefaultInteraction&#34;:false},&#34;isAnonymous&#34;:true,&#34;test&#34;:false,&#34;purposes&#34;:[{&#34;Id&#34;:&#34;C4E88E4C-4719-46DC-8E09-8CC94AE8D850&#34;,&#34;TransactionType&#34;:&#34;NO_CHOICE&#34;},{&#34;Id&#34;:&#34;722AFC4A-AA30-43ED-9F9C-9928F9059071&#34;,&#34;TransactionType&#34;:&#34;CONFIRMED&#34;},{&#34;Id&#34;:&#34;18750554-6B0B-41FE-8C7F-0C7FD1B268F4&#34;,&#34;TransactionType&#34;:&#34;CONFIRMED&#34;},{&#34;Id&#34;:&#34;BA5B7360-F1FB-4224-8FDB-BDF9F48D1B39&#34;,&#34;TransactionType&#34;:&#34;CONFIRMED&#34;}],&#34;dsDataElements&#34;:{&#34;InteractionType&#34;:&#34;Banner - Close&#34;,&#34;Country&#34;:&#34;us&#34;,&#34;UserAgent&#34;:&#34;Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.25 Safari/537.36&#34;,&#34;ConsentModel&#34;:&#34;opt-out&#34;}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/snowflake/analyticsdata?timestamp=1746729080595", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;event-name&#34;:&#34;close cart&#34;,&#34;tag&#34;:&#34;Button&#34;,&#34;tag-location&#34;:&#34;Cart Icon&#34;,&#34;landingpage-url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;window-inner-width&#34;:1280,&#34;window-inner-height&#34;:720,&#34;body-client-width&#34;:1280,&#34;body-client-height&#34;:720,&#34;screen-width&#34;:1280,&#34;screen-height&#34;:720,&#34;screen-width-height&#34;:&#34;1280x720&#34;,&#34;device-pixel-ratio&#34;:1,&#34;browser-name&#34;:&#34;Chrome&#34;,&#34;browser-version&#34;:&#34;136.0.7103.25&#34;,&#34;os-name&#34;:&#34;Windows&#34;,&#34;os-version&#34;:&#34;NT 10.0&#34;,&#34;engine-name&#34;:&#34;Blink&#34;,&#34;engine-version&#34;:&#34;&#34;,&#34;device-type&#34;:&#34;desktop&#34;,&#34;session_id&#34;:&#34;8ffaa7ae-9b97-463d-b437-ce63a0483042&#34;,&#34;productName&#34;:&#34;&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://obs.fishrobotflower.com/mon", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`e=37dfbd8ee84e00126ee8c037e347829d9225c24f567d43d6da1908be6245cad7bd70a976710ce60ed89373bfe70e9c20c1e53e8d5a138c6c2717071a10acf9f29f671a82d589562b6d1ef878215780338a33c103300376c5065359600d5ccfbf394f77be26bb25cb43e2913df05365ac0b297d1bdc57e446f492d7da3dbb280ef67ccaa8073f8d0e3717224793845731f360b3f493a0180dec1edae97dfa2bc8169b1adc597cff3200e714561c4b92177af998ffe4198b6dec06c213f85e162ae7d133722b325f817c99ec59b058609fc6e359143e3dd385293e88864c06513c157a77bb9e70392652b48d1c2ad7f4ec3ee3b8192d4079b4a7a6928677a0dadf5ae8489e5d2e019cbecbf7af2b95dfe57594351ccdeb8b795904fd7367a095166dbdd0aa73d902fd12b0801d91499d9978953b67919d27dc6796d3bbe193fdbd4c38fc28b0bdfd354371fe8f719aa61af7010642dd4245c2979684c0ec8825ce3bc68ebe817fa27d763bc97f27db838003679ea875ac41bfc003d667a87e5346588c6eee5f8b89ee7ef8df78f02da7c19cdb27d60a0a2dea5d540b28073ed7ab67b355c4cb9764fdde25b3d02dd0c2b1be7c85b7f409746c89a91004d7b41d5107f7fae939fa2890f802fdfc3ce52dd36dcc068a50aefc2ecd0823e2ac90633f3d6402340a5b13ace68e65d205be5752517faac71386cd464887eccdb0ae6ddcece3f880ffadf84b4835fe4586a85e6b844a601fd16a59746ef07aa6771ec1f9dd71c83cc52fe0a82cc8a40e99721322210760129cc0e4db207a559a184689d2d9bbb0c24f049eb0be7f8ad2757adbce0f3a68fd8dc45020b23b98599e1cff22981ad8d40efbb337fe868bd49e5b7feee3337f9f71a68aacb6119d6d280b02956a5f4db6a25b93f112ca9887d552cd839b2e083745d321b951ef0ce6936384335d20463e7b78e3538e863a4bcf1cc7c2f03da8cb0be60d76e678072c8df4c3491fb725e78402db1930074eba3c2e8bafd061d52c49e4ff15ebbd9d88e11aa90924605fe06b3716d9c9584eb02ee6e0c02a8bd5eb0c547ca5633e507e867459e6c8e66c40cd260bbbd6c353d5c6f88f0f30b05bd61b781488173175d4cc1615da92b2aabde142a563f930b77096d65d9afff22a2bfdd215808cb4b1c4f6cadf741fddf42f54872dd257c81cf3ee28a8cad1c731efc4c7f2fb5a69cba2b5af51cb5594404921413d88823335dc60f64f8ac965ef0522e2714c9766e69c7575b97521525796f3d12f87c74c603d1e6ac2d583ffce74db&amp;cri=jfECv66HI9&amp;sf=0&amp;dc=d397Y3MEY3R0LnRjdHRjdAVjdHQoKShrJCopJS0vKCFjdHRjdAVyd3FjdAVyd3V2Y3QFd3F3cmN0BXd%2Bdn9jdAV3c39wY3QFd352f2N0BXdzd3VjdAV2Y3QFcnJjdAV3f3V2Y3QFd391dmNzAmBxd3tjcQRjdHQ1Y3R0Y3UHd2N0BWN0dCtjdHRjdQd3Y3QFY3R0MWN0dGN1B3djdAVjdHQkGWN0dGN1B3d1cmNxAmB3cnd7Y3EEY3R0NCM1Y3R0Y3UHY3R0Y3QFdmN0BXdjdAV0Y3QFdWN0BXJjdAVzY3QFcGN0BXFjdAV%2BY3QFf2N0BSdjdAUkY3QFJWN0BSJjdAUjY3QFIGN0BSFjdAUuY3QFL2N0BSxjdAUtY3QFKmN0BStjdAUoY3QFKWN0BTZjdAU3Y3QFNGN0BTVjdAUyY3QFM2N0BTBjdAUxY3QFPmN0BT9jdAU8Y3QFY3B2Y3QFY3MFY3MFY3QFY3QFY3MEY3QFY3MCY3QFY3QFY3QFY3UCY3QFYwV0YwdxY3QFY3QFY3QFa2N0BWhjdAVhY3QFY3UEY3QFY3QAY3R0Y3QFY3R0JBljdHRjdQd3Y3ECYHdzf3tjcQRjdHQlLic0IS8oIWN0dGN1ByAnKjUjY3QFY3R0JS4nNCEvKCESLysjY3R0Y3UHY3R0GWN0dGN0BWN0dCIvNSUuJzQhLyghEi8rI2N0dGN1B3V2cHZ2Y3QFY3R0KiMwIypjdHRjdQd2aH9%2BY3QFY3R0JBljdHRjdQd2Y3ECYHV0cHtjcwRjcQRjdHQkY3R0Y3UHdGN0BWN0dDVjdHRjdQdjdHR3Y3R0Y3ECY3QFY3EEY3R0JGN0dGN1B3VjdAVjdHQ1Y3R0Y3UHY3R0d2N0dGNxAmNzAmB1cHd7Y3EEY3R0NWN0dGN1B3djdAVjdHQwY3R0Y3UHdH9xfn5%2Ff3V2dmN0BWN0dCowY3R0Y3UHY3MEY3MEY3R0FScrJygyLidjdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV3Y3MCY3QFY3MEY3R0Byc0KShjdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0ByokIzQyY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dAcqLyUjY3R0Y3QFY3R0a2N0dGN0BWN0dC8yaw8SY3R0Y3QFdmNzAmN0BWNzBGN0dAcqMCdjdHRjdAVjdHRrY3R0Y3QFY3R0NTBrFQNjdHRjdAV2Y3MCY3QFY3MEY3R0BytjBXVjB38qLyNjdHRjdAVjdHRrY3R0Y3QFY3R0IDRrBQdjdHRjdAV2Y3MCY3QFY3MEY3R0BysvNCdjdHRjdAVjdHRrY3R0Y3QFY3R0KzVrCx9jdHRjdAV2Y3MCY3QFY3MEY3R0BygoJ2N0dGN0BWN0dGtjdHRjdAVjdHQiI2sCA2N0dGN0BXZjcwJjdAVjcwRjdHQHNDIuMzRjdHRjdAVjdHRrY3R0Y3QFY3R0IyhrAQRjdHRjdAV2Y3MCY3QFY3MEY3R0BCciY3R2CCMxNWN0dGN0BWN0dGtjdHRjdAVjdHQjKGsTFWN0dGN0BXZjcwJjdAVjcwRjdHQEJy4uY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dAQjKio1Y3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dAQpLyghY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dAQzJCQqIzVjdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0BSc0Ky8yY3R0Y3QFY3R0a2N0dGN0BWN0dC4jaw8KY3R0Y3QFdmNzAmN0BWNzBGN0dAUnMi4jNC8oI2N0dGN0BWN0dGtjdHRjdAVjdHQjKGsHE2N0dGN0BXZjcwJjdAVjcwRjdHQFIyoqKTVjdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0AicrJz8nKDIvY3R0Y3QFY3R0a2N0dGN0BWN0dC8iaw8CY3R0Y3QFdmNzAmN0BWNzBGN0dAInKC8jKmN0dm4ZdGN0dm4Zcm9vY3R0Y3QFY3R0a2N0dGN0BWN0dCMoawEEY3R0Y3QFdmNzAmN0BWNzBGN0dAInKC8jKmN0dm4ANCMoJS5jdHZuADQnKCUjb29jdHRjdAVjdHRrY3R0Y3QFY3R0IDRrABRjdHRjdAV2Y3MCY3QFY3MEY3R0Aic0LydjdHRjdAVjdHRrY3R0Y3QFY3R0JCFrBAFjdHRjdAV2Y3MCY3QFY3MEY3R0AyIiP2N0dm4BIzQrJyhjdHZuASM0KycoP29vY3R0Y3QFY3R0a2N0dGN0BWN0dCIjawIDY3R0Y3QFdmNzAmN0BWNzBGN0dAMiIj9jdHZuGXRjdHZuGXJvb2N0dGN0BWN0dGtjdHRjdAVjdHQjKGsBBGN0dGN0BXZjcwJjdAVjcwRjdHQDIiI%2FY3R2bhl0Y3R2bhl1b29jdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0AyIiP2N0dm4VNicoLzUuY3R2bhU2Jy8ob29jdHRjdAVjdHRrY3R0Y3QFY3R0IzVrAxVjdHRjdAV2Y3MCY3QFY3MEY3R0AyIiP2N0dm4VNicoLzUuY3R2bgsjPi8lKW9vY3R0Y3QFY3R0a2N0dGN0BWN0dCM1awseY3R0Y3QFdmNzAmN0BWNzBGN0dAMiIj9jdHZuAC8oKC81LmN0dm4ALygqJygib29jdHRjdAVjdHRrY3R0Y3QFY3R0IC9rAA9jdHRjdAV2Y3MCY3QFY3MEY3R0AyIiP2N0dm4ANCMoJS5jdHZuBScoJyInb29jdHRjdAVjdHRrY3R0Y3QFY3R0IDRrBQdjdHRjdAV2Y3MCY3QFY3MEY3R0AyIiP2N0dm4ANCMoJS5jdHZuADQnKCUjb29jdHRjdAVjdHRrY3R0Y3QFY3R0IDRrABRjdHRjdAV2Y3MCY3QFY3MEY3R0AyIiP2N0dm4PMicqLycoY3R2bg8yJyo%2Fb29jdHRjdAVjdHRrY3R0Y3QFY3R0LzJrDxJjdHRjdAV2Y3MCY3QFY3MEY3R0AyIiP2N0dm4MJzYnKCM1I2N0dm4MJzYnKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dCwnawwWY3R0Y3QFdmNzAmN0BWNzBGN0dAMiIj9jdHZuDSk0IycoY3R2bhUpMzIuY3R2DSk0Iydvb2N0dGN0BWN0dGtjdHRjdAVjdHQtKWsNFGN0dGN0BXZjcwJjdAVjcwRjdHQDIiI%2FY3R2bhYpNDIzITMjNSNjdHZuBDQnPC8qb29jdHRjdAVjdHRrY3R0Y3QFY3R0NjJrBBRjdHRjdAV2Y3MCY3QFY3MEY3R0AyIiP2N0dm4FLi8oIzUjY3R2bgUuLygnY3R2KycvKConKCJvb2N0dGN0BWN0dGtjdHRjdAVjdHQ8LmsFCGN0dGN0BXZjcwJjdAVjcwRjdHQDIiI%2FY3R2bgUuLygjNSNjdHZuEicvMScob29jdHRjdAVjdHRrY3R0Y3QFY3R0PC5rEhFjdHRjdAV2Y3MCY3QFY3MEY3R0AyoqIyhjdHRjdAVjdHRrY3R0Y3QFY3R0KCprBANjdHRjdAV2Y3MCY3QFY3MEY3R0ACopY3R2bgEjNCsnKGN0dm4BIzQrJyg%2Fb29jdHRjdAVjdHRrY3R0Y3QFY3R0IiNrAgNjdHRjdAV2Y3MCY3QFY3MEY3R0ACopY3R2bhl0Y3R2bhlyb29jdHRjdAVjdHRrY3R0Y3QFY3R0IyhrAQRjdHRjdAV2Y3MCY3QFY3MEY3R0ACopY3R2bhl0Y3R2bhl1b29jdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0ACopY3R2bhU2JygvNS5jdHZuFTYnLyhvb2N0dGN0BWN0dGtjdHRjdAVjdHQjNWsDFWN0dGN0BXZjcwJjdAVjcwRjdHQAKiljdHZuFTYnKC81LmN0dm4LIz4vJSlvb2N0dGN0BWN0dGtjdHRjdAVjdHQjNWsLHmN0dGN0BXZjcwJjdAVjcwRjdHQAKiljdHZuAC8oKC81LmN0dm4ALygqJygib29jdHRjdAVjdHRrY3R0Y3QFY3R0IC9rAA9jdHRjdAV2Y3MCY3QFY3MEY3R0ACopY3R2bgA0IyglLmN0dm4FJygnIidvb2N0dGN0BWN0dGtjdHRjdAVjdHQgNGsFB2N0dGN0BXZjcwJjdAVjcwRjdHQAKiljdHZuADQjKCUuY3R2bgA0JyglI29vY3R0Y3QFY3R0a2N0dGN0BWN0dCA0awAUY3R0Y3QFdmNzAmN0BWNzBGN0dAAqKWN0dm4PMicqLycoY3R2bg8yJyo%2Fb29jdHRjdAVjdHRrY3R0Y3QFY3R0LzJrDxJjdHRjdAV2Y3MCY3QFY3MEY3R0ACopY3R2bgwnNicoIzUjY3R2bgwnNicob29jdHRjdAVjdHRrY3R0Y3QFY3R0LCdrDBZjdHRjdAV2Y3MCY3QFY3MEY3R0ACopY3R2bg0pNCMnKGN0dm4VKTMyLmN0dg0pNCMnb29jdHRjdAVjdHRrY3R0Y3QFY3R0LSlrDRRjdHRjdAV2Y3MCY3QFY3MEY3R0ACopY3R2bhYpNDIzITMjNSNjdHZuBDQnPC8qb29jdHRjdAVjdHRrY3R0Y3QFY3R0NjJrBBRjdHRjdAV2Y3MCY3QFY3MEY3R0ACopY3R2bgUuLygjNSNjdHZuBS4vKCdjdHYrJy8oKicoIm9vY3R0Y3QFY3R0a2N0dGN0BWN0dDwuawUIY3R0Y3QFdmNzAmN0BWNzBGN0dAAqKWN0dm4FLi8oIzUjY3R2bhInLzEnKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dDwuaxIRY3R0Y3QFdmNzAmN0BWNzBGN0dAA0IyJjdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0ASkpImN0dggjMTVjdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0ASk0IikoY3R0Y3QFY3R0a2N0dGN0BWN0dCMoawcTY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiKydjdHZuASM0KycoY3R2bgEjNCsnKD9vb2N0dGN0BWN0dGtjdHRjdAVjdHQiI2sCA2N0dGN0BXZjcwJjdAVjcwRjdHQBNCcoIisnY3R2bhl0Y3R2bhlyb29jdHRjdAVjdHRrY3R0Y3QFY3R0IyhrAQRjdHRjdAV2Y3MCY3QFY3MEY3R0ATQnKCIrJ2N0dm4ZdGN0dm4ZdW9vY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiKydjdHZuFTYnKC81LmN0dm4VNicvKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dCM1awMVY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiKydjdHZuFTYnKC81LmN0dm4LIz4vJSlvb2N0dGN0BWN0dGtjdHRjdAVjdHQjNWsLHmN0dGN0BXZjcwJjdAVjcwRjdHQBNCcoIisnY3R2bgAvKCgvNS5jdHZuAC8oKicoIm9vY3R0Y3QFY3R0a2N0dGN0BWN0dCAvawAPY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiKydjdHZuADQjKCUuY3R2bgUnKCciJ29vY3R0Y3QFY3R0a2N0dGN0BWN0dCA0awUHY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiKydjdHZuADQjKCUuY3R2bgA0JyglI29vY3R0Y3QFY3R0a2N0dGN0BWN0dCA0awAUY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiKydjdHZuDzInKi8nKGN0dm4PMicqP29vY3R0Y3QFY3R0a2N0dGN0BWN0dC8yaw8SY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiKydjdHZuDCc2JygjNSNjdHZuDCc2Jyhvb2N0dGN0BWN0dGtjdHRjdAVjdHQsJ2sMFmN0dGN0BXZjcwJjdAVjcwRjdHQBNCcoIisnY3R2bg0pNCMnKGN0dm4VKTMyLmN0dg0pNCMnb29jdHRjdAVjdHRrY3R0Y3QFY3R0LSlrDRRjdHRjdAV2Y3MCY3QFY3MEY3R0ATQnKCIrJ2N0dm4WKTQyMyEzIzUjY3R2bgQ0JzwvKm9vY3R0Y3QFY3R0a2N0dGN0BWN0dDYyawQUY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiKydjdHZuBS4vKCM1I2N0dm4FLi8oJ2N0disnLygqJygib29jdHRjdAVjdHRrY3R0Y3QFY3R0PC5rBQhjdHRjdAV2Y3MCY3QFY3MEY3R0ATQnKCIrJ2N0dm4FLi8oIzUjY3R2bhInLzEnKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dDwuaxIRY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiNidjdHZuASM0KycoY3R2bgEjNCsnKD9vb2N0dGN0BWN0dGtjdHRjdAVjdHQiI2sCA2N0dGN0BXZjcwJjdAVjcwRjdHQBNCcoIjYnY3R2bhl0Y3R2bhlyb29jdHRjdAVjdHRrY3R0Y3QFY3R0IyhrAQRjdHRjdAV2Y3MCY3QFY3MEY3R0ATQnKCI2J2N0dm4ZdGN0dm4ZdW9vY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiNidjdHZuFTYnKC81LmN0dm4VNicvKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dCM1awMVY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiNidjdHZuFTYnKC81LmN0dm4LIz4vJSlvb2N0dGN0BWN0dGtjdHRjdAVjdHQjNWsLHmN0dGN0BXZjcwJjdAVjcwRjdHQBNCcoIjYnY3R2bgAvKCgvNS5jdHZuAC8oKicoIm9vY3R0Y3QFY3R0a2N0dGN0BWN0dCAvawAPY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiNidjdHZuADQjKCUuY3R2bgUnKCciJ29vY3R0Y3QFY3R0a2N0dGN0BWN0dCA0awUHY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiNidjdHZuADQjKCUuY3R2bgA0JyglI29vY3R0Y3QFY3R0a2N0dGN0BWN0dCA0awAUY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiNidjdHZuDzInKi8nKGN0dm4PMicqP29vY3R0Y3QFY3R0a2N0dGN0BWN0dC8yaw8SY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiNidjdHZuDCc2JygjNSNjdHZuDCc2Jyhvb2N0dGN0BWN0dGtjdHRjdAVjdHQsJ2sMFmN0dGN0BXZjcwJjdAVjcwRjdHQBNCcoIjYnY3R2bg0pNCMnKGN0dm4VKTMyLmN0dg0pNCMnb29jdHRjdAVjdHRrY3R0Y3QFY3R0LSlrDRRjdHRjdAV2Y3MCY3QFY3MEY3R0ATQnKCI2J2N0dm4WKTQyMyEzIzUjY3R2bgQ0JzwvKm9vY3R0Y3QFY3R0a2N0dGN0BWN0dDYyawQUY3R0Y3QFdmNzAmN0BWNzBGN0dAE0JygiNidjdHZuBS4vKCM1I2N0dm4FLi8oJ2N0disnLygqJygib29jdHRjdAVjdHRrY3R0Y3QFY3R0PC5rBQhjdHRjdAV2Y3MCY3QFY3MEY3R0ATQnKCI2J2N0dm4FLi8oIzUjY3R2bhInLzEnKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dDwuaxIRY3R0Y3QFdmNzAmN0BWNzBGN0dA4nMjIpNC9jdHRjdAVjdHRrY3R0Y3QFY3R0LCdrDBZjdHRjdAV2Y3MCY3QFY3MEY3R0DiMqIygnY3R0Y3QFY3R0a2N0dGN0BWN0dCIjawIDY3R0Y3QFdmNzAmN0BWNzBGN0dA8pJygnY3R0Y3QFY3R0a2N0dGN0BWN0dDQpaxQJY3R0Y3QFdmNzAmN0BWNzBGN0dAwnJTczIzVjdHRjdAVjdHRrY3R0Y3QFY3R0IDRrABRjdHRjdAV2Y3MCY3QFY3MEY3R0DCM1MiM0Y3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dAwpJygnY3R0Y3QFY3R0a2N0dGN0BWN0dDYyaxYSY3R0Y3QFdmNzAmN0BWNzBGN0dAwzKC8pNGN0dGN0BWN0dGtjdHRjdAVjdHQjKGsTFWN0dGN0BXZjcwJjdAVjcwRjdHQNJyg%2FJ2N0dGN0BWN0dGtjdHRjdAVjdHQyLmsSDmN0dGN0BXZjcwJjdAVjcwRjdHQNJzQjKGN0dGN0BWN0dGtjdHRjdAVjdHQjKGsHE2N0dGN0BXZjcwJjdAVjcwRjdHQNJzIuP2N0dGN0BWN0dGtjdHRjdAVjdHQjKGsTFWN0dGN0BXZjcwJjdAVjcwRjdHQNPyktKWN0dGN0BWN0dGtjdHRjdAVjdHQsJ2sMFmN0dGN0BXZjcwJjdAVjcwRjdHQKJygnY3R0Y3QFY3R0a2N0dGN0BWN0dC40aw4UY3R0Y3QFdmNzAmN0BWNzBGN0dAonMzQnY3R0Y3QFY3R0a2N0dGN0BWN0dDUtaxUNY3R0Y3QFdmNzAmN0BWNzBGN0dAojLS4nY3R0Y3QFY3R0a2N0dGN0BWN0dC4vaw8IY3R0Y3QFdmNzAmN0BWNzBGN0dAojNT8nY3R0Y3QFY3R0a2N0dGN0BWN0dDMtaxMHY3R0Y3QFdmNzAmN0BWNzBGN0dAovawszY3R0Y3QFY3R0a2N0dGN0BWN0dDwuawUIY3R0Y3QFdmNzAmN0BWNzBGN0dAovKC5jdHRjdAVjdHRrY3R0Y3QFY3R0MC9rEAhjdHRjdAV2Y3MCY3QFY3MEY3R0CjMlLycoJ2N0dGN0BWN0dGtjdHRjdAVjdHQ2MmsEFGN0dGN0BXZjcwJjdAVjcwRjdHQLJywjImN0dGN0BWN0dGtjdHRjdAVjdHQnNGt2dndjdHRjdAV2Y3MCY3QFY3MEY3R0Cyc0LyNjdHRjdAVjdHRrY3R0Y3QFY3R0IDRrABRjdHRjdAV2Y3MCY3QFY3MEY3R0Cyc0Mi4nY3R0Y3QFY3R0a2N0dGN0BWN0dCMoawEEY3R0Y3QFdmNzAmN0BWNzBGN0dAsnNDIvKGN0dGN0BWN0dGtjdHRjdAVjdHQiI2sCA2N0dGN0BXZjcwJjdAVjcwRjdHQLIy8sLydjdHRjdAVjdHRrY3R0Y3QFY3R0PC5rEhFjdHRjdAV2Y3MCY3QFY3MEY3R0CyMqLygnY3R0Y3QFY3R0a2N0dGN0BWN0dCMqawEUY3R0Y3QFdmNzAmN0BWNzBGN0dAsvKiMoJ2N0dGN0BWN0dGtjdHRjdAVjdHQ0M2sUE2N0dGN0BXZjcwJjdAVjcwRjdHQLKS80J2N0dGN0BWN0dGtjdHRjdAVjdHQjKGsPA2N0dGN0BXZjcwJjdAVjcwRjdHQLKSgyNSNjdHRjdAVjdHRrY3R0Y3QFY3R0JSdrAxVjdHRjdAV2Y3MCY3QFY3MEY3R0C2MFdWMEdSgvJSdjdHRjdAVjdHRrY3R0Y3QFY3R0IzVrAxVjdHRjdAV2Y3MCY3QFY3MEY3R0CC8lLT9jdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0CCk0J2N0dGN0BWN0dGtjdHRjdAVjdHQoJGsICWN0dGN0BXZjcwJjdAVjcwRjdHQJaxQjKGN0dGN0BWN0dGtjdHRjdAVjdHQsJ2sMFmN0dGN0BXZjcwJjdAVjcwRjdHQJNCEnKGN0dGN0BWN0dGtjdHRjdAVjdHQjKGsTFWN0dGN0BXZjcwJjdAVjcwRjdHQWJzMqLygnY3R0Y3QFY3R0a2N0dGN0BWN0dCM1awseY3R0Y3QFdmNzAmN0BWNzBGN0dBQnKjYuY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dBQjIyJjdHZuASM0KycoY3R2bgEjNCsnKD9vb2N0dGN0BWN0dGtjdHRjdAVjdHQiI2sCA2N0dGN0BXZjcwJjdAVjcwRjdHQUIyMiY3R2bhl0Y3R2bhlyb29jdHRjdAVjdHRrY3R0Y3QFY3R0IyhrAQRjdHRjdAV2Y3MCY3QFY3MEY3R0FCMjImN0dm4ZdGN0dm4ZdW9vY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dBQjIyJjdHZuFTYnKC81LmN0dm4VNicvKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dCM1awMVY3R0Y3QFdmNzAmN0BWNzBGN0dBQjIyJjdHZuFTYnKC81LmN0dm4LIz4vJSlvb2N0dGN0BWN0dGtjdHRjdAVjdHQjNWsLHmN0dGN0BXZjcwJjdAVjcwRjdHQUIyMiY3R2bgAvKCgvNS5jdHZuAC8oKicoIm9vY3R0Y3QFY3R0a2N0dGN0BWN0dCAvawAPY3R0Y3QFdmNzAmN0BWNzBGN0dBQjIyJjdHZuADQjKCUuY3R2bgUnKCciJ29vY3R0Y3QFY3R0a2N0dGN0BWN0dCA0awUHY3R0Y3QFdmNzAmN0BWNzBGN0dBQjIyJjdHZuDzInKi8nKGN0dm4PMicqP29vY3R0Y3QFY3R0a2N0dGN0BWN0dC8yaw8SY3R0Y3QFdmNzAmN0BWNzBGN0dBQjIyJjdHZuDCc2JygjNSNjdHZuDCc2Jyhvb2N0dGN0BWN0dGtjdHRjdAVjdHQsJ2sMFmN0dGN0BXZjcwJjdAVjcwRjdHQUIyMiY3R2bg0pNCMnKGN0dm4VKTMyLmN0dg0pNCMnb29jdHRjdAVjdHRrY3R0Y3QFY3R0LSlrDRRjdHRjdAV2Y3MCY3QFY3MEY3R0FCMjImN0dm4WKTQyMyEzIzUjY3R2bgQ0JzwvKm9vY3R0Y3QFY3R0a2N0dGN0BWN0dDYyawQUY3R0Y3QFdmNzAmN0BWNzBGN0dBQjIyJjdHZuBS4vKCM1I2N0dm4FLi8oJ2N0disnLygqJygib29jdHRjdAVjdHRrY3R0Y3QFY3R0PC5rBQhjdHRjdAV2Y3MCY3QFY3MEY3R0FCMjImN0dm4FLi8oIzUjY3R2bhInLzEnKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dDwuaxIRY3R0Y3QFdmNzAmN0BWNzBGN0dBQvNS4vY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaw8IY3R0Y3QFdmNzAmN0BWNzBGN0dBQpJS0pY3R2bgEjNCsnKGN0dm4BIzQrJyg%2Fb29jdHRjdAVjdHRrY3R0Y3QFY3R0IiNrAgNjdHRjdAV2Y3MCY3QFY3MEY3R0FCklLSljdHZuGXRjdHZuGXJvb2N0dGN0BWN0dGtjdHRjdAVjdHQjKGsBBGN0dGN0BXZjcwJjdAVjcwRjdHQUKSUtKWN0dm4ZdGN0dm4ZdW9vY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dBQpJS0pY3R2bhU2JygvNS5jdHZuFTYnLyhvb2N0dGN0BWN0dGtjdHRjdAVjdHQjNWsDFWN0dGN0BXZjcwJjdAVjcwRjdHQUKSUtKWN0dm4VNicoLzUuY3R2bgsjPi8lKW9vY3R0Y3QFY3R0a2N0dGN0BWN0dCM1awseY3R0Y3QFdmNzAmN0BWNzBGN0dBQpJS0pY3R2bgAvKCgvNS5jdHZuAC8oKicoIm9vY3R0Y3QFY3R0a2N0dGN0BWN0dCAvawAPY3R0Y3QFdmNzAmN0BWNzBGN0dBQpJS0pY3R2bgA0IyglLmN0dm4FJygnIidvb2N0dGN0BWN0dGtjdHRjdAVjdHQgNGsFB2N0dGN0BXZjcwJjdAVjcwRjdHQUKSUtKWN0dm4ANCMoJS5jdHZuADQnKCUjb29jdHRjdAVjdHRrY3R0Y3QFY3R0IDRrABRjdHRjdAV2Y3MCY3QFY3MEY3R0FCklLSljdHZuDzInKi8nKGN0dm4PMicqP29vY3R0Y3QFY3R0a2N0dGN0BWN0dC8yaw8SY3R0Y3QFdmNzAmN0BWNzBGN0dBQpJS0pY3R2bgwnNicoIzUjY3R2bgwnNicob29jdHRjdAVjdHRrY3R0Y3QFY3R0LCdrDBZjdHRjdAV2Y3MCY3QFY3MEY3R0FCklLSljdHZuDSk0IycoY3R2bhUpMzIuY3R2DSk0Iydvb2N0dGN0BWN0dGtjdHRjdAVjdHQtKWsNFGN0dGN0BXZjcwJjdAVjcwRjdHQUKSUtKWN0dm4WKTQyMyEzIzUjY3R2bgQ0JzwvKm9vY3R0Y3QFY3R0a2N0dGN0BWN0dDYyawQUY3R0Y3QFdmNzAmN0BWNzBGN0dBQpJS0pY3R2bgUuLygjNSNjdHZuBS4vKCdjdHYrJy8oKicoIm9vY3R0Y3QFY3R0a2N0dGN0BWN0dDwuawUIY3R0Y3QFdmNzAmN0BWNzBGN0dBQpJS0pY3R2bgUuLygjNSNjdHZuEicvMScob29jdHRjdAVjdHRrY3R0Y3QFY3R0PC5rEhFjdHRjdAV2Y3MCY3QFY3MEY3R0FScoIj9jdHZuASM0KycoY3R2bgEjNCsnKD9vb2N0dGN0BWN0dGtjdHRjdAVjdHQiI2sCA2N0dGN0BXZjcwJjdAVjcwRjdHQVJygiP2N0dm4ZdGN0dm4Zcm9vY3R0Y3QFY3R0a2N0dGN0BWN0dCMoawEEY3R0Y3QFdmNzAmN0BWNzBGN0dBUnKCI%2FY3R2bhl0Y3R2bhl1b29jdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0FScoIj9jdHZuFTYnKC81LmN0dm4VNicvKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dCM1awMVY3R0Y3QFdmNzAmN0BWNzBGN0dBUnKCI%2FY3R2bhU2JygvNS5jdHZuCyM%2BLyUpb29jdHRjdAVjdHRrY3R0Y3QFY3R0IzVrCx5jdHRjdAV2Y3MCY3QFY3MEY3R0FScoIj9jdHZuAC8oKC81LmN0dm4ALygqJygib29jdHRjdAVjdHRrY3R0Y3QFY3R0IC9rAA9jdHRjdAV2Y3MCY3QFY3MEY3R0FScoIj9jdHZuADQjKCUuY3R2bgUnKCciJ29vY3R0Y3QFY3R0a2N0dGN0BWN0dCA0awUHY3R0Y3QFdmNzAmN0BWNzBGN0dBUnKCI%2FY3R2bgA0IyglLmN0dm4ANCcoJSNvb2N0dGN0BWN0dGtjdHRjdAVjdHQgNGsAFGN0dGN0BXZjcwJjdAVjcwRjdHQVJygiP2N0dm4PMicqLycoY3R2bg8yJyo%2Fb29jdHRjdAVjdHRrY3R0Y3QFY3R0LzJrDxJjdHRjdAV2Y3MCY3QFY3MEY3R0FScoIj9jdHZuDCc2JygjNSNjdHZuDCc2Jyhvb2N0dGN0BWN0dGtjdHRjdAVjdHQsJ2sMFmN0dGN0BXZjcwJjdAVjcwRjdHQVJygiP2N0dm4NKTQjJyhjdHZuFSkzMi5jdHYNKTQjJ29vY3R0Y3QFY3R0a2N0dGN0BWN0dC0paw0UY3R0Y3QFdmNzAmN0BWNzBGN0dBUnKCI%2FY3R2bhYpNDIzITMjNSNjdHZuBDQnPC8qb29jdHRjdAVjdHRrY3R0Y3QFY3R0NjJrBBRjdHRjdAV2Y3MCY3QFY3MEY3R0FScoIj9jdHZuBS4vKCM1I2N0dm4FLi8oJ2N0disnLygqJygib29jdHRjdAVjdHRrY3R0Y3QFY3R0PC5rBQhjdHRjdAV2Y3MCY3QFY3MEY3R0FScoIj9jdHZuBS4vKCM1I2N0dm4SJy8xJyhvb2N0dGN0BWN0dGtjdHRjdAVjdHQ8LmsSEWN0dGN0BXZjcwJjdAVjcwRjdHQVJzQnY3R0Y3QFY3R0a2N0dGN0BWN0dCInawINY3R0Y3QFdmNzAmN0BWNzBGN0dBUnMjNjdHRjdAVjdHRrY3R0Y3QFY3R0IC9rAA9jdHRjdAV2Y3MCY3QFY3MEY3R0FS4jKiojP2N0dm4BIzQrJyhjdHZuASM0KycoP29vY3R0Y3QFY3R0a2N0dGN0BWN0dCIjawIDY3R0Y3QFdmNzAmN0BWNzBGN0dBUuIyoqIz9jdHZuGXRjdHZuGXJvb2N0dGN0BWN0dGtjdHRjdAVjdHQjKGsBBGN0dGN0BXZjcwJjdAVjcwRjdHQVLiMqKiM%2FY3R2bhl0Y3R2bhl1b29jdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0FS4jKiojP2N0dm4VNicoLzUuY3R2bhU2Jy8ob29jdHRjdAVjdHRrY3R0Y3QFY3R0IzVrAxVjdHRjdAV2Y3MCY3QFY3MEY3R0FS4jKiojP2N0dm4VNicoLzUuY3R2bgsjPi8lKW9vY3R0Y3QFY3R0a2N0dGN0BWN0dCM1awseY3R0Y3QFdmNzAmN0BWNzBGN0dBUuIyoqIz9jdHZuAC8oKC81LmN0dm4ALygqJygib29jdHRjdAVjdHRrY3R0Y3QFY3R0IC9rAA9jdHRjdAV2Y3MCY3QFY3MEY3R0FS4jKiojP2N0dm4ANCMoJS5jdHZuBScoJyInb29jdHRjdAVjdHRrY3R0Y3QFY3R0IDRrBQdjdHRjdAV2Y3MCY3QFY3MEY3R0FS4jKiojP2N0dm4ANCMoJS5jdHZuADQnKCUjb29jdHRjdAVjdHRrY3R0Y3QFY3R0IDRrABRjdHRjdAV2Y3MCY3QFY3MEY3R0FS4jKiojP2N0dm4PMicqLycoY3R2bg8yJyo%2Fb29jdHRjdAVjdHRrY3R0Y3QFY3R0LzJrDxJjdHRjdAV2Y3MCY3QFY3MEY3R0FS4jKiojP2N0dm4MJzYnKCM1I2N0dm4MJzYnKG9vY3R0Y3QFY3R0a2N0dGN0BWN0dCwnawwWY3R0Y3QFdmNzAmN0BWNzBGN0dBUuIyoqIz9jdHZuDSk0IycoY3R2bhUpMzIuY3R2DSk0Iydvb2N0dGN0BWN0dGtjdHRjdAVjdHQtKWsNFGN0dGN0BXZjcwJjdAVjcwRjdHQVLiMqKiM%2FY3R2bhYpNDIzITMjNSNjdHZuBDQnPC8qb29jdHRjdAVjdHRrY3R0Y3QFY3R0NjJrBBRjdHRjdAV2Y3MCY3QFY3MEY3R0FS4jKiojP2N0dm4FLi8oIzUjY3R2bgUuLygnY3R2KycvKConKCJvb2N0dGN0BWN0dGtjdHRjdAVjdHQ8LmsFCGN0dGN0BXZjcwJjdAVjcwRjdHQVLiMqKiM%2FY3R2bgUuLygjNSNjdHZuEicvMScob29jdHRjdAVjdHRrY3R0Y3QFY3R0PC5rEhFjdHRjdAV2Y3MCY3QFY3MEY3R0FS8oLC9jdHRjdAVjdHRrY3R0Y3QFY3R0PC5rDg1jdHRjdAV2Y3MCY3QFY3MEY3R0FTM2IzQ1Mic0Y3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dBIjNTUnY3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxwHY3R0Y3QFdmNzAmN0BWNzBGN0dBIuKSsnNWN0dGN0BWN0dGtjdHRjdAVjdHQgNGsAFGN0dGN0BXZjcwJjdAVjcwRjdHQSLygnY3R0Y3QFY3R0a2N0dGN0BWN0dDUqaxUPY3R0Y3QFdmNzAmN0BWNzBGN0dBIvKCEyLyghY3R0Y3QFY3R0a2N0dGN0BWN0dDwuawUIY3R0Y3QFdmNzAmN0BWNzBGN0dBI0LygpLyI1Y3R0Y3QFY3R0a2N0dGN0BWN0dCMoaxMVY3R0Y3QFdmNzAmN0BWNzBGN0dBJjBXVjBAUoIiNjdHRjdAVjdHRrY3R0Y3QFY3R0LjNrDhNjdHRjdAV2Y3MCY3QFY3MEY3R0ES4vNTYjNGN0dGN0BWN0dGtjdHRjdAVjdHQjKGsTFWN0dGN0BXZjcwJjdAVjcwRjdHQRKSQkKiNjdHRjdAVjdHRrY3R0Y3QFY3R0IyhrExVjdHRjdAV2Y3MCY3QFY3MEY3R0HicoIiM0Y3R0Y3QFY3R0a2N0dGN0BWN0dCgqawgKY3R0Y3QFdmNzAmN0BWNzBGN0dB8jKiInY3R0Y3QFY3R0a2N0dGN0BWN0dDI0axIUY3R0Y3QFdmNzAmN0BWNzBGN0dB8zazUuM2N0dGN0BWN0dGtjdHRjdAVjdHQ8LmsFCGN0dGN0BXZjcwJjdAVjcwRjdHQfMygnY3R0Y3QFY3R0a2N0dGN0BWN0dC0paw0UY3R0Y3QFdmNzAmN0BWNzBGN0dBwnNDApPmN0dGN0BWN0dGtjdHRjdAVjdHQjKGsTFWN0dGN0BXZjcwJjdAVjcwRjdHQcKTUvJ2N0dGN0BWN0dGtjdHRjdAVjdHQ2KmsWCmN0dGN0BXZjcwJjdAVjcwRjdHQcMzwnKCdjdHRjdAVjdHRrY3R0Y3QFY3R0JTVrBRxjdHRjdAV2Y3MCY3MCY3QFY3R0MCM0Y3R0Y3UHdGN0BWN0dCQZY3R0Y3UHd2NxAmBzdHZ7Y3EEY3R0NmN0dGN1B2N0dAsnJQ8oMiMqY3R0Y3QFY3R0KmN0dGN1B2NzBGN0dCMoaxMVY3R0Y3MCY3QFY3R0IitjdHRjdQd%2BY3QFY3R0LiVjdHRjdQd3dGN0BWN0dDMnImN0dGN1B2NxBGN0dCc0JS4vMiMlMjM0I2N0dGN1B2N0dD5wcmN0dGN0BWN0dCQvMigjNTVjdHRjdQdjdHRwcmN0dGN0BWN0dCQ0JygiNWN0dGN1B2NzBGNxBGN0dCQ0JygiY3R0Y3UHY3R0BS40KSsvMytjdHRjdAVjdHQwIzQ1LykoY3R0Y3UHY3R0d3VwY3R0Y3ECY3QFY3EEY3R0JDQnKCJjdHRjdQdjdHQOIyciKiM1NQUuNCkrI2N0dGN0BWN0dDAjNDUvKShjdHRjdQdjdHR3dXBjdHRjcQJjdAVjcQRjdHQkNCcoImN0dGN1B2N0dAgpMmgHY3QABDQnKCJjdHRjdAVjdHQwIzQ1LykoY3R0Y3UHY3R0f39jdHRjcQJjcwJjdAVjdHQgMyoqECM0NS8pKAovNTJjdHRjdQdjcwRjcQRjdHQkNCcoImN0dGN1B2N0dAUuNCkrLzMrY3R0Y3QFY3R0MCM0NS8pKGN0dGN1B2N0dHd1cGh2aHF3dnVodHNjdHRjcQJjdAVjcQRjdHQkNCcoImN0dGN1B2N0dA4jJyIqIzU1BS40KSsjY3R0Y3QFY3R0MCM0NS8pKGN0dGN1B2N0dHd1cGh2aHF3dnVodHNjdHRjcQJjdAVjcQRjdHQkNCcoImN0dGN1B2N0dAgpMmgHY3QABDQnKCJjdHRjdAVjdHQwIzQ1LykoY3R0Y3UHY3R0f39odmh2aHZjdHRjcQJjcwJjdAVjdHQrKSQvKiNjdHRjdQcgJyo1I2N0BWN0dCspIiMqY3R0Y3UHY3R0Y3R0Y3QFY3R0NionMiApNCtjdHRjdQdjdHQRLygiKTE1Y3R0Y3QFY3R0NionMiApNCsQIzQ1LykoY3R0Y3UHY3R0d3ZodmN0dGN0BWN0dDMnADMqKhAjNDUvKShjdHRjdQdjdHR3dXBodmhxd3Z1aHRzY3R0Y3QFY3R0MSkxcHJjdHRjdQcgJyo1I2NxAmN0BWN0dCQZY3R0Y3UHc2NxAmB%2BcXJ7Y3EEY3R0NWN0dGN1B3VjdAVjdHQzMDYnY3R0Y3UHdmN0BWN0dCQZY3R0Y3UHd2NxAg%3D%3D&amp;cp=c&amp;gtm=WyJPbmVUcnVzdExvYWRlZCIsIk9wdGFub25Mb2FkZWQiLCJPbmVUcnVzdEdyb3Vwc1VwZGF0ZWQiLCJjb252ZXJzaW9uIl0%3D&amp;gac=462250241.1746729079&amp;mm=1%7C1115%2C518%2C0%2C0%2C1115%2C518%2C518%2C1%2C1115%2C518%7C1115%2C518%2C0%2C0%2C1115%2C518%2C518%2C1%2C1115%2C518&amp;md=1%7C1115%2C518%2C1115%2C518%2C520%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C1115%2C518%2C1115%2C518%2C520%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518&amp;mu=1%7C1115%2C518%2C1115%2C518%2C521%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C1115%2C518%2C1115%2C518%2C521%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518&amp;cl=1%7C1115%2C518%2C1115%2C518%2C528%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C1115%2C518%2C1115%2C518%2C528%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518&amp;tb=1&amp;gi=9%2C2448%7C12%2C2449%7C10%2C2450%7C14%2C2452%7C11%2C2452%7C13%2C2452%7C21%2C2461&amp;ws=1280x720&amp;wos=1280x720&amp;ver=13&amp;fi=517&amp;ti=530&amp;mo=0&amp;pn=2461&amp;spn=1931&amp;fp=516`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/runtime~main.9754bbac.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/10.f16292bd.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/main~493df0b3.9fda564a.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/runtime~main.9754bbac.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/10.f16292bd.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/main~493df0b3.9fda564a.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC71c39d3f3a504282b82da0f673925a9c-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC700b4b859fd0473ca58c1fad9bc2bf9f-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCdbc6d4c5412646698048b64684bb035d-source.min.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://zndnxlcj0ulh6d1zq-crowdstrike.siteintercept.qualtrics.com/WRSiteInterceptEngine/?Q_ZID=ZN_dnXlCJ0uLH6d1ZQ", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/54.1ade363e.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/38.ef717b79.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/23.60057654.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/20.2ffef383.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/45.3e7e52c2.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/21.b3438b1b.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/27.3951aad8.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/16.44924e69.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/12.d33926cb.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/19.8e79a39a.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/52.df339939.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/43.ebd6caf4.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/30.57dfb56c.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/44.5b845402.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/22.4cb40074.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/9.6ac3976b.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/9.a767dca8.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/17.22abfce0.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/17.90ff250d.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/25.87fbc779.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/18.c419c295.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/54.1ade363e.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/38.ef717b79.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/23.60057654.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/20.2ffef383.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/45.3e7e52c2.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/21.b3438b1b.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/27.3951aad8.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/16.44924e69.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/12.d33926cb.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/19.8e79a39a.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/52.df339939.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/43.ebd6caf4.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/30.57dfb56c.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/44.5b845402.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/22.4cb40074.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/9.6ac3976b.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/9.a767dca8.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/17.22abfce0.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/17.90ff250d.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/25.87fbc779.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/18.c419c295.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://px.ads.linkedin.com/wa/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;pids&#34;:[64444],&#34;scriptVersion&#34;:213,&#34;time&#34;:1746729080668,&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;url&#34;:&#34;https://crowdstrike.com/en-us/&#34;,&#34;pageTitle&#34;:&#34;CrowdStrike: We Stop Breaches with AI-native Cybersecurity&#34;,&#34;websiteSignalRequestId&#34;:&#34;7546bc93-8490-0577-6ba5-0d773e0cdee4&#34;,&#34;isTranslated&#34;:false,&#34;liFatId&#34;:&#34;&#34;,&#34;liGiant&#34;:&#34;&#34;,&#34;misc&#34;:{&#34;psbState&#34;:-4},&#34;isLinkedInApp&#34;:false,&#34;hem&#34;:null,&#34;signalType&#34;:&#34;CLICK&#34;,&#34;href&#34;:&#34;&#34;,&#34;domAttributes&#34;:{&#34;elementSemanticType&#34;:&#34;ICON&#34;,&#34;elementValue&#34;:null,&#34;elementType&#34;:null,&#34;tagName&#34;:&#34;BUTTON&#34;,&#34;backgroundImageSrc&#34;:&#34;https://cdn.cookielaw.org/logos/static/ot_close.svg&#34;,&#34;imageSrc&#34;:null,&#34;imageAlt&#34;:null,&#34;innerText&#34;:&#34;&#34;,&#34;elementTitle&#34;:&#34;Close&#34;,&#34;cursor&#34;:&#34;pointer&#34;},&#34;innerElements&#34;:null,&#34;elementCrumbsTree&#34;:[{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:15,&#34;id&#34;:&#34;onetrust-consent-sdk&#34;},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:1,&#34;id&#34;:&#34;onetrust-banner-sdk&#34;,&#34;classes&#34;:[&#34;bottom&#34;,&#34;otFlat&#34;,&#34;vertical-align-content&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:1,&#34;id&#34;:&#34;onetrust-close-btn-container&#34;},{&#34;tagName&#34;:&#34;button&#34;,&#34;nthChild&#34;:0,&#34;classes&#34;:[&#34;banner-close-button&#34;,&#34;onetrust-close-btn-handler&#34;,&#34;onetrust-close-btn-ui&#34;,&#34;ot-close-icon&#34;]}],&#34;isFilteredByClient&#34;:false}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/snowflake/analyticsdata?timestamp=1746729080671", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;event-name&#34;:&#34;close cart&#34;,&#34;tag&#34;:&#34;Button&#34;,&#34;tag-location&#34;:&#34;Cart Icon&#34;,&#34;landingpage-url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;window-inner-width&#34;:1280,&#34;window-inner-height&#34;:720,&#34;body-client-width&#34;:1280,&#34;body-client-height&#34;:720,&#34;screen-width&#34;:1280,&#34;screen-height&#34;:720,&#34;screen-width-height&#34;:&#34;1280x720&#34;,&#34;device-pixel-ratio&#34;:1,&#34;browser-name&#34;:&#34;Chrome&#34;,&#34;browser-version&#34;:&#34;136.0.7103.25&#34;,&#34;os-name&#34;:&#34;Windows&#34;,&#34;os-version&#34;:&#34;NT 10.0&#34;,&#34;engine-name&#34;:&#34;Blink&#34;,&#34;engine-version&#34;:&#34;&#34;,&#34;device-type&#34;:&#34;desktop&#34;,&#34;session_id&#34;:&#34;8ffaa7ae-9b97-463d-b437-ce63a0483042&#34;,&#34;productName&#34;:&#34;&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.zi-scripts.com/unified/v1/master/getSubscriptions", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.partnerstack.com/v1/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.zi-scripts.com/zi-tag.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 304 {
		t.Errorf("Expected status %d, got %d", 304, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://collector-20290.tvsquared.com/tv2track.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/trw?aid=crowdstrike&amp;amp;trwv.uid=crowdstrike-1746729079514-1f81ceec&amp;amp;trwv.vc=1&amp;amp;trwsa.sid=crowdstrike-1746729079515-66d02f45&amp;amp;trwsb.cpv=2&amp;amp;ctzo=-07:00&amp;amp;uri=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;pm=&amp;amp;viewedTypes=&amp;amp;rts=1746729080686", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://alb.reddit.com/rp.gif?ts=1746729080687&amp;amp;id=t2_2n40s6z5&amp;amp;event=PageVisit&amp;amp;m.itemCount=&amp;amp;m.value=&amp;amp;m.valueDecimal=&amp;amp;m.currency=&amp;amp;m.transactionId=&amp;amp;m.customEventName=&amp;amp;m.products=&amp;amp;m.conversionId=&amp;amp;uuid=2bc44b75-14f8-4772-a1c1-5d39c30b5042&amp;amp;aaid=&amp;amp;em=&amp;amp;pn=&amp;amp;external_id=&amp;amp;idfa=&amp;amp;integration=reddit&amp;amp;partner=&amp;amp;opt_out=0&amp;amp;sh=1280&amp;amp;sw=720&amp;amp;v=rdt_00917691&amp;amp;dpm=&amp;amp;dpcc=&amp;amp;dprc=", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/trw?aid=crowdstrike&amp;amp;trwv.uid=crowdstrike-1746729079514-1f81ceec&amp;amp;trwv.vc=1&amp;amp;trwsa.sid=crowdstrike-1746729079515-66d02f45&amp;amp;trwsb.cpv=3&amp;amp;ctzo=-07:00&amp;amp;uri=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;pm=&amp;amp;viewedTypes=&amp;amp;rts=1746729080691", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://alb.reddit.com/rp.gif?ts=1746729080692&amp;amp;id=t2_2n40s6z5&amp;amp;event=PageVisit&amp;amp;m.itemCount=&amp;amp;m.value=&amp;amp;m.valueDecimal=&amp;amp;m.currency=&amp;amp;m.transactionId=&amp;amp;m.customEventName=&amp;amp;m.products=&amp;amp;m.conversionId=&amp;amp;uuid=2bc44b75-14f8-4772-a1c1-5d39c30b5042&amp;amp;aaid=&amp;amp;em=&amp;amp;pn=&amp;amp;external_id=&amp;amp;idfa=&amp;amp;integration=reddit&amp;amp;partner=&amp;amp;opt_out=0&amp;amp;sh=1280&amp;amp;sw=720&amp;amp;v=rdt_00917691&amp;amp;dpm=&amp;amp;dpcc=&amp;amp;dprc=", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ibc-flow.techtarget.com/a/gif.gif?actTypeId=31&amp;amp;cid=3218843&amp;amp;r=1746729080698&amp;amp;ref=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;version=2.4", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != -1 {
		t.Errorf("Expected status %d, got %d", -1, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ibc-flow.techtarget.com/a/gif.gif?actTypeId=31&amp;amp;cid=3218843&amp;amp;r=1746729080702&amp;amp;ref=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;version=2.4", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != -1 {
		t.Errorf("Expected status %d, got %d", -1, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://c.contentsquare.net/errors?v=15.85.1&amp;amp;pid=29632&amp;amp;pn=1&amp;amp;sn=1&amp;amp;uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;amp;ct=0", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;errors&#34;:[{&#34;errorType&#34;:&#34;jsError&#34;,&#34;message&#34;:&#34;Script error.&#34;,&#34;filename&#34;:&#34;&#34;,&#34;lineno&#34;:0,&#34;colno&#34;:0,&#34;pageUrl&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;rt&#34;:989}]}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=3e40d144-374d-442e-a9ff-4d2d60049d2a&amp;amp;batch_time=1746729080722", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080283,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;26d58040-2072-4252-920e-ce68d229e094&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://bat.bing.com/action/0?ti=163002607&amp;Ver=2&amp;mid=ae9bdae8-4702-4299-b1a6-34555b231a31&amp;bo=2&amp;sid=a2fec0b02c3a11f095697fbf91a88d6a&amp;vid=a2fec0802c3a11f0b5c3b9fd0fc28d87&amp;vids=0&amp;msclkid=N&amp;ec=CHEQ&amp;el=Invalid_Users&amp;ev=0&amp;ea=Invalid_Users&amp;en=Y&amp;p=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;sw=1280&amp;sh=720&amp;sc=24&amp;nwd=1&amp;evt=custom&amp;cdb=AQET&amp;rn=414587&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:58000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080076,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;919b9ed3-dd1d-489c-99a7-c01d30501748&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://www.googleadservices.com/pagead/conversion/797629828/?label=hozuCPn52LoYEIS7q_wC&amp;guid=ON&amp;script=0&#34;,&#34;protocol&#34;:&#34;h3&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:277500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:42,&#34;encoded_body_size&#34;:42,&#34;decoded_body_size&#34;:42,&#34;transfer_size&#34;:342,&#34;download&#34;:{&#34;duration&#34;:200000,&#34;start&#34;:277300000},&#34;first_byte&#34;:{&#34;duration&#34;:59300000,&#34;start&#34;:218000000},&#34;redirect&#34;:{&#34;duration&#34;:217600000,&#34;start&#34;:0}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080073,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;10784070-24bd-4868-9482-630336171a3f&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://www.googleadservices.com/pagead/conversion/797629828/?random=1746729080069&amp;cv=11&amp;fst=1746729080069&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gcl_ctr=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;label=hozuCPn52LoYEIS7q_wC&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;gtm_ee=1&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;capi=1&amp;data=event%3Dconversion&amp;rfmt=3&amp;fmt=4&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:291800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:6193,&#34;encoded_body_size&#34;:2391,&#34;decoded_body_size&#34;:6193,&#34;transfer_size&#34;:2691,&#34;download&#34;:{&#34;duration&#34;:700000,&#34;start&#34;:291100000},&#34;first_byte&#34;:{&#34;duration&#34;:187100000,&#34;start&#34;:104000000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080283,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;693073e1-e834-4c8b-b024-730385bdf0c8&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://bat.bing.com/action/0?ti=163002607&amp;Ver=2&amp;mid=ae9bdae8-4702-4299-b1a6-34555b231a31&amp;bo=1&amp;sid=a2fec0b02c3a11f095697fbf91a88d6a&amp;vid=a2fec0802c3a11f0b5c3b9fd0fc28d87&amp;vids=1&amp;msclkid=N&amp;uach=pv%3D10.0&amp;pi=0&amp;lg=en-US&amp;sw=1280&amp;sh=720&amp;sc=24&amp;nwd=1&amp;tl=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;p=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;r=&amp;lt=569&amp;evt=pageLoad&amp;sv=1&amp;cdb=AQET&amp;rn=206258&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:124000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080365,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;52b818e9-1d53-498a-9159-b952b4844ed6&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://googleads.g.doubleclick.net/pagead/viewthroughconversion/797629828/?random=160478099&amp;cv=11&amp;fst=1746729080069&amp;bg=ffffff&amp;guid=ON&amp;async=1&amp;gcl_ctr=1&amp;gtm=45je5570h1v894068940za200&amp;gcd=13l3l3l3l1l1&amp;dma=0&amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;u_w=1280&amp;u_h=720&amp;url=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;label=hozuCPn52LoYEIS7q_wC&amp;hn=www.googleadservices.com&amp;frm=0&amp;tiba=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;did=dYWJhMj&amp;gdid=dYWJhMj&amp;gtm_ee=1&amp;npa=0&amp;pscdl=noapi&amp;auid=1475816545.1746729079&amp;uaa=x64&amp;uab=64&amp;uafvl=Chromium%3B136.0.7103.25%7CHeadlessChrome%3B136.0.7103.25%7CNot.A%252FBrand%3B99.0.0.0&amp;uamb=0&amp;uam=&amp;uap=Windows&amp;uapv=10.0&amp;uaw=0&amp;fledge=1&amp;capi=1&amp;data=event%3Dconversion&amp;fmt=3&amp;ct_cookie_present=false&amp;crd=CPLOsQIIobixAgixwbECCLDBsQIIscOxAgiKxbECCMLJsQIItMaxAgiQybECCNPFsQII68yxAgjPzrECCP7OsQII1c-xAkoVdHJpZ2dlciwgZXZlbnQtc291cmNlWgMKAQFiBAoCAgM&amp;pscrd=IhMI34Lk8sCUjQMVqKOOCB2_0BUQMgwIA2IICAAQABgAIAAyDAgEYggIABAAGAAgADIMCAdiCAgAEAAYACAAMgwICGIICAAQABgAIAAyDAgJYggIABAAGAAgADIMCApiCAgAEAAYACAAMgwIAmIICAAQABgAIAAyDAgLYggIABAAGAAgADIMCBViCAgAEAAYACAAMgwIH2IICAAQABgAIAAyDAgTYggIABAAGAAgADIMCBJiCAgAEAAYACAAOhxodHRwczovL3d3dy5jcm93ZHN0cmlrZS5jb20vQlZDaEVJOEt6eHdBWVFvTE9vdXAtQXU5bjlBUklyQVBSSTZXMURUUl94MUZMVUdqZUNFdGtHbWhGY0RkQmdnRHdlMVQ4SUNUYVpZRTdHN095bFlqY0ZBUQ&#34;,&#34;protocol&#34;:&#34;h3&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:153300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:42,&#34;encoded_body_size&#34;:42,&#34;decoded_body_size&#34;:42,&#34;transfer_size&#34;:342,&#34;download&#34;:{&#34;duration&#34;:300000,&#34;start&#34;:153000000},&#34;first_byte&#34;:{&#34;duration&#34;:71800000,&#34;start&#34;:81200000},&#34;redirect&#34;:{&#34;duration&#34;:80700000,&#34;start&#34;:0}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080590,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d2d5489b-8a17-4092-a3c1-8ccad0607cd1&#34;,&#34;type&#34;:&#34;beacon&#34;,&#34;url&#34;:&#34;https://www.facebook.com/tr/&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:31100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080586,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8e3a806b-7f47-4114-8bcc-b34f2a5b9726&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:66000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:204,&#34;url&#34;:&#34;https://px.ads.linkedin.com/wa/&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080636,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1dcb800d-fc61-477a-98a9-f8182e4d951c&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC71c39d3f3a504282b82da0f673925a9c-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:20000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:351,&#34;encoded_body_size&#34;:224,&#34;decoded_body_size&#34;:351,&#34;transfer_size&#34;:524,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:19900000},&#34;first_byte&#34;:{&#34;duration&#34;:19700000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080540,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1bbe87e1-fb06-4632-bfcc-2eef5b55353b&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:120000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/ga/sgm?sid=crowdstrike-1746729079515-66d02f45&amp;1746729080540&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080540,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;ce4246dd-88ed-4120-abd6-64df79429f8c&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:125000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/rtp/api/v1_1/visitor?sid=crowdstrike-1746729079515-66d02f45&amp;aid=crowdstrike&amp;1746729080540&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0}},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080704,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;,&#34;in_foreground&#34;:true},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;error&#34;:{&#34;id&#34;:&#34;3f9b27ce-584a-46e2-b2d8-baca98cbd25d&#34;,&#34;message&#34;:&#34;Uncaught \&#34;Script error.\&#34;&#34;,&#34;source&#34;:&#34;source&#34;,&#34;stack&#34;:&#34;Error: Script error.\n  at undefined @ &#34;,&#34;handling&#34;:&#34;unhandled&#34;,&#34;source_type&#34;:&#34;browser&#34;},&#34;type&#34;:&#34;error&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080534,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;77288571-cdb2-4ec8-8398-902398b02e1f&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://js.zi-scripts.com/zi-tag.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:148600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080657,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1eded538-01b9-4dbe-98d9-244d1ddac992&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RC700b4b859fd0473ca58c1fad9bc2bf9f-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:27100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:1203,&#34;encoded_body_size&#34;:477,&#34;decoded_body_size&#34;:1203,&#34;transfer_size&#34;:777,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:27000000},&#34;first_byte&#34;:{&#34;duration&#34;:26700000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080657,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b687c7ec-c67f-49d6-b328-6f4ddbc3b687&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://assets.adobedtm.com/d72cd986aea0/f7467a554824/80aec2017ba3/RCdbc6d4c5412646698048b64684bb035d-source.min.js&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:26800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:571,&#34;encoded_body_size&#34;:337,&#34;decoded_body_size&#34;:571,&#34;transfer_size&#34;:637,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:26800000},&#34;first_byte&#34;:{&#34;duration&#34;:26600000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/40.eeb001f3.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/40.3c4347cc.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/msg?a=2&amp;amp;sid=crowdstrike-1746729079515-66d02f45&amp;amp;aid=crowdstrike&amp;amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;viewedTypes=&amp;amp;0.4334833343192306&amp;amp;rts=1746729080753", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/msg?a=2&amp;amp;sid=crowdstrike-1746729079515-66d02f45&amp;amp;aid=crowdstrike&amp;amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;viewedTypes=&amp;amp;0.8025280330938003&amp;amp;rts=1746729080753", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/0.0b2ebd4a.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/33.0e6e41b3.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/28.b5e8f5e1.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/28.d0035c49.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/26.c695453b.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/26.e987b304.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.zi-scripts.com/unified/v1/master/getSubscriptions", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.zi-scripts.com/unified/v1/master/getSubscriptions", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/widgetapp/2025-05-08-12-13-34/remediation/remediation_1746706414131.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/remediations/consolidated/2376540/VuzDL5zIiHUEAGOy.json", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/styles/2025-05-08-12-13-34/widget_base.css?v=1746706414131", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/styles/2025-05-08-12-13-34/widget_base.css?v=1746706414131", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/styles/2025-05-08-12-13-34/widget_base.css?v=1746706414131", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/0.0b2ebd4a.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/4.07aa08a5.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/4.6c355058.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/1.d2d44206.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/1.c4d0069a.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/3.e276c1eb.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/29.812d5a7c.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/29.6ae501ce.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://grsm.io/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23?get_pscd=true", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://grsm.io/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23?get_pscd=true", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://bootstrap.driftapi.com/widget_bootstrap/ping/v2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`ping_context=%7B%22embedId%22%3A%229d4udx6ceimp%22%2C%22pageUrl%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F%22%7D`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/msg?a=2&amp;amp;sid=crowdstrike-1746729079515-66d02f45&amp;amp;aid=crowdstrike&amp;amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;viewedTypes=&amp;amp;0.8068380355532484&amp;amp;rts=1746729080805", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://sjrtp1.marketo.com/gw1/msg?a=2&amp;amp;sid=crowdstrike-1746729079515-66d02f45&amp;amp;aid=crowdstrike&amp;amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;viewedTypes=&amp;amp;0.588399501303551&amp;amp;rts=1746729080809", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "*/*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/widgetapp/images/body_wh.svg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/widgetapp/images/spin_wh.svg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/remediation/2025-05-08-12-13-34/paid/remediation-tool.js?ts=1746706414131", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://siteintercept.qualtrics.com/dxjsmodule/8.b8b768ed34bc7505f919.chunk.js?Q_CLIENTVERSION=2.28.0&amp;amp;Q_CLIENTTYPE=web&amp;amp;Q_BRANDID=www.crowdstrike.com", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://edge.adobedc.net/ee/or2/v1/privacy/set-consent?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=b8a8506a-974d-4da9-a84a-22f63ae140b3", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;consent&#34;:[{&#34;standard&#34;:&#34;Adobe&#34;,&#34;version&#34;:&#34;1.0&#34;,&#34;value&#34;:{&#34;general&#34;:&#34;in&#34;}}],&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true,&#34;entries&#34;:[{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_identity&#34;,&#34;value&#34;:&#34;CiYyODE4MDkyODQ4NjE4NzUyODY1MDM5MjUzNTg4NzAxMjI5MTU4OFISCJ2jqonrMhABGAEqA09SMjAA8AGdo6qJ6zI=&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_cluster&#34;,&#34;value&#34;:&#34;or2&#34;}]}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=07c1fab2-eb8d-4f87-b25d-5277f0f829a3&amp;amp;batch_time=1746729080834", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080596,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;cbf66936-9940-4628-97e3-1ebc5f485489&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:114000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://obs.fishrobotflower.com/mon&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080668,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;05ab2203-3c11-40b3-817f-f40871124650&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:54000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:204,&#34;url&#34;:&#34;https://px.ads.linkedin.com/wa/&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080596,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;81c85443-2e63-4475-8f36-da975fb8e0b6&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:157200000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/snowflake/analyticsdata?timestamp=1746729080595&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:153,&#34;encoded_body_size&#34;:127,&#34;decoded_body_size&#34;:153,&#34;transfer_size&#34;:427,&#34;download&#34;:{&#34;duration&#34;:600000,&#34;start&#34;:156600000},&#34;first_byte&#34;:{&#34;duration&#34;:156400000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080685,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c545038b-eb3d-4875-b253-304d37b35c46&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://js.zi-scripts.com/zi-tag.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:89300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080685,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;dba424ba-3fc0-40b8-87ef-95b507ab628a&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://js.partnerstack.com/v1/&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:106500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080524,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;2a2c14e7-f5f7-4ab8-8e28-cff3697cc6c9&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://js.driftt.com/core?d=1&amp;embedId=9d4udx6ceimp&amp;eId=9d4udx6ceimp&amp;region=US&amp;forceShow=false&amp;skipCampaigns=false&amp;sessionId=bb076d85-176d-40bf-8edd-e4d09c550682&amp;sessionStarted=1746729080.521&amp;campaignRefreshToken=f7d1fcdd-1c5b-41be-9318-104dbcaa8709&amp;hideController=false&amp;pageLoadStartTime=1746729078294&amp;mode=CHAT&amp;driftEnableLog=false&amp;secureIframe=false&amp;u=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:271700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080693,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;4ffea0b6-a833-4d30-9adc-6d5c0652345f&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://alb.reddit.com/rp.gif?ts=1746729080692&amp;id=t2_2n40s6z5&amp;event=PageVisit&amp;m.itemCount=&amp;m.value=&amp;m.valueDecimal=&amp;m.currency=&amp;m.transactionId=&amp;m.customEventName=&amp;m.products=&amp;m.conversionId=&amp;uuid=2bc44b75-14f8-4772-a1c1-5d39c30b5042&amp;aaid=&amp;em=&amp;pn=&amp;external_id=&amp;idfa=&amp;integration=reddit&amp;partner=&amp;opt_out=0&amp;sh=1280&amp;sw=720&amp;v=rdt_00917691&amp;dpm=&amp;dpcc=&amp;dprc=&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:105500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080688,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;d408d327-37f6-4d1b-920d-668be537afa9&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://alb.reddit.com/rp.gif?ts=1746729080687&amp;id=t2_2n40s6z5&amp;event=PageVisit&amp;m.itemCount=&amp;m.value=&amp;m.valueDecimal=&amp;m.currency=&amp;m.transactionId=&amp;m.customEventName=&amp;m.products=&amp;m.conversionId=&amp;uuid=2bc44b75-14f8-4772-a1c1-5d39c30b5042&amp;aaid=&amp;em=&amp;pn=&amp;external_id=&amp;idfa=&amp;integration=reddit&amp;partner=&amp;opt_out=0&amp;sh=1280&amp;sw=720&amp;v=rdt_00917691&amp;dpm=&amp;dpcc=&amp;dprc=&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:112300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080687,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;aaa13dd0-8cd7-4422-941a-be88fb1f683b&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/trw?aid=crowdstrike&amp;trwv.uid=crowdstrike-1746729079514-1f81ceec&amp;trwv.vc=1&amp;trwsa.sid=crowdstrike-1746729079515-66d02f45&amp;trwsb.cpv=2&amp;ctzo=-07:00&amp;uri=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;pm=&amp;viewedTypes=&amp;rts=1746729080686&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:118400000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080692,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6ec7d794-d964-49ae-986b-f6eefe863a92&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/trw?aid=crowdstrike&amp;trwv.uid=crowdstrike-1746729079514-1f81ceec&amp;trwv.vc=1&amp;trwsa.sid=crowdstrike-1746729079515-66d02f45&amp;trwsb.cpv=3&amp;ctzo=-07:00&amp;uri=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;pm=&amp;viewedTypes=&amp;rts=1746729080691&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:117100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080526,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5c397036-6fec-465f-b854-ffe96c440a01&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://js.driftt.com/core/chat?d=1&amp;region=US&amp;driftEnableLog=false&amp;pageLoadStartTime=1746729078294&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:296200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080782,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;48aba59e-c681-4cbb-bbe7-dc1bf55dc44c&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.userway.org/widgetapp/2025-05-08-12-13-34/remediation/remediation_1746706414131.js&#34;,&#34;status_code&#34;:200,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:35700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080783,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;02646c03-fc37-43e4-81cb-aa62707d610b&#34;,&#34;type&#34;:&#34;css&#34;,&#34;url&#34;:&#34;https://cdn.userway.org/styles/2025-05-08-12-13-34/widget_base.css?v=1746706414131&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:37600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080672,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8d88d195-e784-4a40-9c08-070bb6026b08&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:146200000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/snowflake/analyticsdata?timestamp=1746729080671&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:153,&#34;encoded_body_size&#34;:127,&#34;decoded_body_size&#34;:153,&#34;transfer_size&#34;:427,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:146200000},&#34;first_byte&#34;:{&#34;duration&#34;:146000000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080658,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;618b3e4d-1632-4fa7-a242-61d4a72141ac&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://zndnxlcj0ulh6d1zq-crowdstrike.siteintercept.qualtrics.com/WRSiteInterceptEngine/?Q_ZID=ZN_dnXlCJ0uLH6d1ZQ&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:167700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:8825,&#34;encoded_body_size&#34;:3820,&#34;decoded_body_size&#34;:8825,&#34;transfer_size&#34;:4120,&#34;download&#34;:{&#34;duration&#34;:3300000,&#34;start&#34;:164400000},&#34;first_byte&#34;:{&#34;duration&#34;:70700000,&#34;start&#34;:93700000},&#34;connect&#34;:{&#34;duration&#34;:68700000,&#34;start&#34;:24900000},&#34;ssl&#34;:{&#34;duration&#34;:39000000,&#34;start&#34;:54600000},&#34;dns&#34;:{&#34;duration&#34;:24800000,&#34;start&#34;:100000}},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://siteintercept.qualtrics.com/WRSiteInterceptEngine/Targeting.php?Q_ZoneID=ZN_dnXlCJ0uLH6d1ZQ&amp;amp;Q_CLIENTVERSION=2.28.0&amp;amp;Q_CLIENTTYPE=webAdobeLaunch", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`Q_LOC=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;hasAnalyticsConsent=null&amp;Timestamp=2025-05-08T18%3A31%3A20.880Z`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.userway.org/remediations/consolidated/2376540/VuzDL5zIiHUEAGOy.json", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ws.zoominfo.com/pixel/61b22df2e97826001a6d4b6e/?iszitag=true", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://get.crowdstrike.com/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://partnerlinks.io/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://get.crowdstrike.com/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://partnerlinks.io/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ws.zoominfo.com/pixel/61b22df2e97826001a6d4b6e/?iszitag=true", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://metrics.api.drift.com/monitoring/metrics/widget/init/v3", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`init_event=%7B%22eventName%22%3A%22%5BV2%5D%20-%20init%22%2C%22embedId%22%3A%229d4udx6ceimp%22%2C%22sessionId%22%3A%22bb076d85-176d-40bf-8edd-e4d09c550682%22%2C%22instanceId%22%3A%225e2d5498-aa99-4851-a408-bdc0dce124ee%22%2C%22context%22%3A%7B%22url%22%3A%22crowdstrike.com%2Fen-us%22%2C%22hostname%22%3A%22www.crowdstrike.com%22%2C%22timezone%22%3A%22America%2FLos_Angeles%22%2C%22userAgent%22%3A%22Mozilla%2F5.0%20(Windows%20NT%2010.0%3B%20Win64%3B%20x64)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Chrome%2F136.0.7103.25%20Safari%2F537.36%22%2C%22timeToInit%22%3A2712%7D%7D`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://bootstrap.driftapi.com/widget_bootstrap", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`embed_id=9d4udx6ceimp&amp;client_id=f6zuizdyhxrm7r&amp;consent_id=has_consent&amp;lead_id=3a55016c-317c-4ceb-935c-33bce08ec640&amp;marketo_cookie=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;integrations=%7B%22marketo%22%3A%22id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61%22%7D&amp;sessionStarted=1746729080&amp;targeting_context=%7B%22sessionId%22%3A%22f7d1fcdd-1c5b-41be-9318-104dbcaa8709%22%2C%22pageUrl%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F%22%2C%22hasHadConversations%22%3Afalse%2C%22followingCampaignIds%22%3A%5B%5D%2C%22excludedCampaignIds%22%3A%5B%5D%2C%22siteVisits%22%3A1%7D&amp;slimCampaigns=true`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=272e9f18-390c-40d1-8b9c-53348fc2f4be&amp;amp;batch_time=1746729081008", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080754,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;192cd028-7f9a-4182-b1d1-ad6ae6f1c2cc&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/msg?a=2&amp;sid=crowdstrike-1746729079515-66d02f45&amp;aid=crowdstrike&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;viewedTypes=&amp;0.4334833343192306&amp;rts=1746729080753&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:71700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080754,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9c996b45-c6c8-45d0-bce0-6c7c41b668a4&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/msg?a=2&amp;sid=crowdstrike-1746729079515-66d02f45&amp;aid=crowdstrike&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;viewedTypes=&amp;0.8025280330938003&amp;rts=1746729080753&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:95300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080828,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;68c3613b-611d-49da-8bb6-90fed9c78091&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.userway.org/widgetapp/images/body_wh.svg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:33600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080828,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;ab28ff41-12fc-46e8-bbdb-6ec5ad6a3dfc&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.userway.org/widgetapp/images/spin_wh.svg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:34500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080686,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;4e22abfa-dd40-43f8-b81b-ad1d731a3240&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://collector-20290.tvsquared.com/tv2track.js&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:185300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080830,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;72a4ea64-b6cc-4aef-995a-615aa9d66e49&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://siteintercept.qualtrics.com/dxjsmodule/8.b8b768ed34bc7505f919.chunk.js?Q_CLIENTVERSION=2.28.0&amp;Q_CLIENTTYPE=web&amp;Q_BRANDID=www.crowdstrike.com&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:47900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:81168,&#34;encoded_body_size&#34;:22897,&#34;decoded_body_size&#34;:81168,&#34;transfer_size&#34;:23197,&#34;download&#34;:{&#34;duration&#34;:400000,&#34;start&#34;:47500000},&#34;first_byte&#34;:{&#34;duration&#34;:45600000,&#34;start&#34;:1900000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080593,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c86bb390-f12f-4d2a-aa2b-5789440f4351&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:287000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://privacyportal.onetrust.com/request/v1/consentreceipts&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080782,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c9de3330-0e59-479c-83ee-ddcc0caf58d6&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:105000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://cdn.userway.org/remediations/consolidated/2376540/VuzDL5zIiHUEAGOy.json&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080829,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;7e619d58-0560-446e-bdc6-53aeaf469517&#34;,&#34;type&#34;:&#34;js&#34;,&#34;url&#34;:&#34;https://cdn.userway.org/remediation/2025-05-08-12-13-34/paid/remediation-tool.js?ts=1746706414131&#34;,&#34;status_code&#34;:200,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:63800000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080897,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;3af2d7c2-fca1-48e0-a365-1fb89f1bb514&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:6000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://cdn.userway.org/remediations/consolidated/2376540/VuzDL5zIiHUEAGOy.json&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080806,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9c00ada0-86ab-451c-9e34-9a5138736da8&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/msg?a=2&amp;sid=crowdstrike-1746729079515-66d02f45&amp;aid=crowdstrike&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;viewedTypes=&amp;0.8068380355532484&amp;rts=1746729080805&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:100300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080809,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;83549e2a-6f70-4b36-af3f-15f534ae6cd1&#34;,&#34;type&#34;:&#34;other&#34;,&#34;url&#34;:&#34;https://sjrtp1.marketo.com/gw1/msg?a=2&amp;sid=crowdstrike-1746729079515-66d02f45&amp;aid=crowdstrike&amp;ma=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;viewedTypes=&amp;0.588399501303551&amp;rts=1746729080809&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:97700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080683,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;708db27f-d699-4aef-a7ab-4973eaca37aa&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:234000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://js.zi-scripts.com/unified/v1/master/getSubscriptions&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080792,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1f9f6d25-c229-4eee-8dd5-7c8767e445b5&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:143000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://grsm.io/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23?get_pscd=true&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080702,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b09a75cd-7cee-4327-8094-7855a516da07&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:282000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:0,&#34;url&#34;:&#34;https://ibc-flow.techtarget.com/a/gif.gif?actTypeId=31&amp;cid=3218843&amp;r=1746729080702&amp;ref=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;version=2.4&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080698,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;fbdfce65-2118-4ac9-9825-a963b4d2fea2&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:306000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:0,&#34;url&#34;:&#34;https://ibc-flow.techtarget.com/a/gif.gif?actTypeId=31&amp;cid=3218843&amp;r=1746729080698&amp;ref=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;version=2.4&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://obs.fishrobotflower.com/mon", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`e=37dfbd8ee84e00126ee8c037e347829d9225c24f567d43d6da1908be6245cad7bd70a976710ce60ed89373bfe70e9c20c1e53e8d5a138c6c2717071a10acf9f29f671a82d589562b6d1ef878215780338a33c103300376c5065359600d5ccfbf394f77be26bb25cb43e2913df05365ac0b297d1bdc57e446f492d7da3dbb280ef67ccaa8073f8d0e3717224793845731f360b3f493a0180dec1edae97dfa2bc8169b1adc597cff3200e714561c4b92177af998ffe4198b6dec06c213f85e162ae7d133722b325f817c99ec59b058609fc6e359143e3dd385293e88864c06513c157a77bb9e70392652b48d1c2ad7f4ec3ee3b8192d4079b4a7a6928677a0dadf5ae8489e5d2e019cbecbf7af2b95dfe57594351ccdeb8b795904fd7367a095166dbdd0aa73d902fd12b0801d91499d9978953b67919d27dc6796d3bbe193fdbd4c38fc28b0bdfd354371fe8f719aa61af7010642dd4245c2979684c0ec8825ce3bc68ebe817fa27d763bc97f27db838003679ea875ac41bfc003d667a87e5346588c6eee5f8b89ee7ef8df78f02da7c19cdb27d60a0a2dea5d540b28073ed7ab67b355c4cb9764fdde25b3d02dd0c2b1be7c85b7f409746c89a91004d7b41d5107f7fae939fa2890f802fdfc3ce52dd36dcc068a50aefc2ecd0823e2ac90633f3d6402340a5b13ace68e65d205be5752517faac71386cd464887eccdb0ae6ddcece3f880ffadf84b4835fe4586a85e6b844a601fd16a59746ef07aa6771ec1f9dd71c83cc52fe0a82cc8a40e99721322210760129cc0e4db207a559a184689d2d9bbb0c24f049eb0be7f8ad2757adbce0f3a68fd8dc45020b23b98599e1cff22981ad8d40efbb337fe868bd49e5b7feee3337f9f71a68aacb6119d6d280b02956a5f4db6a25b93f112ca9887d552cd839b2e083745d321b951ef0ce6936384335d20463e7b78e3538e863a4bcf1cc7c2f03da8cb0be60d76e678072c8df4c3491fb725e78402db1930074eba3c2e8bafd061d52c49e4ff15ebbd9d88e11aa90924605fe06b3716d9c9584eb02ee6e0c02a8bd5eb0c547ca5633e507e867459e6c8e66c40cd260bbbd6c353d5c6f88f0f30b05bd61b781488173175d4cc1615da92b2aabde142a563f930b77096d65d9afff22a2bfdd215808cb4b1c4f6cadf741fddf42f54872dd257c81cf3ee28a8cad1c731efc4c7f2fb5a69cba2b5af51cb5594404921413d88823335dc60f64f8ac965ef0522e2714c9766e69c7575b97521525796f3d12f87c74c603d1e6ac2d583ffce74db&amp;cri=jfECv66HI9&amp;sf=0&amp;dc=&amp;cp=custom&amp;gtm=WyJPbmVUcnVzdExvYWRlZCIsIk9wdGFub25Mb2FkZWQiLCJPbmVUcnVzdEdyb3Vwc1VwZGF0ZWQiLCJjb252ZXJzaW9uIl0%3D&amp;gac=462250241.1746729079&amp;mm=2%7C1115%2C518%2C0%2C0%2C1115%2C518%2C518%2C1%2C1115%2C518%7C1115%2C518%2C0%2C0%2C1115%2C518%2C602%2C1%2C1115%2C518&amp;md=2%7C1115%2C518%2C1115%2C518%2C520%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C1115%2C518%2C1115%2C518%2C603%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518&amp;mu=2%7C1115%2C518%2C1115%2C518%2C521%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C1115%2C518%2C1115%2C518%2C603%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518&amp;cl=2%7C1115%2C518%2C1115%2C518%2C528%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C1115%2C518%2C1115%2C518%2C605%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518&amp;tb=1&amp;gi=9%2C2448%7C12%2C2449%7C10%2C2450%7C14%2C2452%7C11%2C2452%7C13%2C2452%7C21%2C2461%7C18%2C2862&amp;ws=1280x720&amp;wos=1280x720&amp;ver=13&amp;fi=517&amp;ti=1002&amp;mo=0&amp;pn=2934&amp;spn=1931&amp;fp=516&amp;snt=1&amp;sc=5%7C0%2C2%2C931%2C1%7C0%2C49%2C997%2C1`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://obs.fishrobotflower.com/mon", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`e=37dfbd8ee84e00126ee8c037e347829d9225c24f567d43d6da1908be6245cad7bd70a976710ce60ed89373bfe70e9c20c1e53e8d5a138c6c2717071a10acf9f29f671a82d589562b6d1ef878215780338a33c103300376c5065359600d5ccfbf394f77be26bb25cb43e2913df05365ac0b297d1bdc57e446f492d7da3dbb280ef67ccaa8073f8d0e3717224793845731f360b3f493a0180dec1edae97dfa2bc8169b1adc597cff3200e714561c4b92177af998ffe4198b6dec06c213f85e162ae7d133722b325f817c99ec59b058609fc6e359143e3dd385293e88864c06513c157a77bb9e70392652b48d1c2ad7f4ec3ee3b8192d4079b4a7a6928677a0dadf5ae8489e5d2e019cbecbf7af2b95dfe57594351ccdeb8b795904fd7367a095166dbdd0aa73d902fd12b0801d91499d9978953b67919d27dc6796d3bbe193fdbd4c38fc28b0bdfd354371fe8f719aa61af7010642dd4245c2979684c0ec8825ce3bc68ebe817fa27d763bc97f27db838003679ea875ac41bfc003d667a87e5346588c6eee5f8b89ee7ef8df78f02da7c19cdb27d60a0a2dea5d540b28073ed7ab67b355c4cb9764fdde25b3d02dd0c2b1be7c85b7f409746c89a91004d7b41d5107f7fae939fa2890f802fdfc3ce52dd36dcc068a50aefc2ecd0823e2ac90633f3d6402340a5b13ace68e65d205be5752517faac71386cd464887eccdb0ae6ddcece3f880ffadf84b4835fe4586a85e6b844a601fd16a59746ef07aa6771ec1f9dd71c83cc52fe0a82cc8a40e99721322210760129cc0e4db207a559a184689d2d9bbb0c24f049eb0be7f8ad2757adbce0f3a68fd8dc45020b23b98599e1cff22981ad8d40efbb337fe868bd49e5b7feee3337f9f71a68aacb6119d6d280b02956a5f4db6a25b93f112ca9887d552cd839b2e083745d321b951ef0ce6936384335d20463e7b78e3538e863a4bcf1cc7c2f03da8cb0be60d76e678072c8df4c3491fb725e78402db1930074eba3c2e8bafd061d52c49e4ff15ebbd9d88e11aa90924605fe06b3716d9c9584eb02ee6e0c02a8bd5eb0c547ca5633e507e867459e6c8e66c40cd260bbbd6c353d5c6f88f0f30b05bd61b781488173175d4cc1615da92b2aabde142a563f930b77096d65d9afff22a2bfdd215808cb4b1c4f6cadf741fddf42f54872dd257c81cf3ee28a8cad1c731efc4c7f2fb5a69cba2b5af51cb5594404921413d88823335dc60f64f8ac965ef0522e2714c9766e69c7575b97521525796f3d12f87c74c603d1e6ac2d583ffce74db&amp;cri=jfECv66HI9&amp;sf=0&amp;dc=&amp;cp=1&amp;gtm=WyJPbmVUcnVzdExvYWRlZCIsIk9wdGFub25Mb2FkZWQiLCJPbmVUcnVzdEdyb3Vwc1VwZGF0ZWQiLCJjb252ZXJzaW9uIl0%3D&amp;gac=462250241.1746729079&amp;mm=2%7C1115%2C518%2C0%2C0%2C1115%2C518%2C518%2C1%2C1115%2C518%7C1115%2C518%2C0%2C0%2C1115%2C518%2C602%2C1%2C1115%2C518&amp;md=2%7C1115%2C518%2C1115%2C518%2C520%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C1115%2C518%2C1115%2C518%2C603%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518&amp;mu=2%7C1115%2C518%2C1115%2C518%2C521%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C1115%2C518%2C1115%2C518%2C603%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518&amp;cl=2%7C1115%2C518%2C1115%2C518%2C528%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C1115%2C518%2C1115%2C518%2C605%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518&amp;tb=1&amp;gi=9%2C2448%7C12%2C2449%7C10%2C2450%7C14%2C2452%7C11%2C2452%7C13%2C2452%7C21%2C2461%7C18%2C2862&amp;ws=1280x720&amp;wos=1280x720&amp;ver=13&amp;fi=517&amp;ti=1007&amp;mo=0&amp;pn=2938&amp;spn=1931&amp;fp=516&amp;snt=1&amp;sc=5%7C0%2C2%2C931%2C1%7C0%2C49%2C997%2C1`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://ws.zoominfo.com/pixel/61b22df2e97826001a6d4b6e/?iszitag=true", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/ipv?_biz_r=&amp;amp;_biz_h=-1719904874&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729079103&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=0&amp;amp;a=crowdstrike.com&amp;amp;rnd=480761&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729079104", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizibly.com/u?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729079104&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;a=crowdstrike.com&amp;amp;rnd=110587&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729079104", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/u?mapType=mkto&amp;amp;mapValue=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729080105&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=1&amp;amp;a=crowdstrike.com&amp;amp;rnd=598845&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729080105", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/u?mapType=ecid&amp;amp;mapValue=06D71E9261F941560A495CD6%40AdobeOrg_28180928486187528650392535887012291588&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729080106&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=2&amp;amp;a=crowdstrike.com&amp;amp;rnd=228061&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729080208", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://bat.bing.com/action/0?ti=163002607&amp;amp;Ver=2&amp;amp;mid=ae9bdae8-4702-4299-b1a6-34555b231a31&amp;amp;bo=1&amp;amp;sid=a2fec0b02c3a11f095697fbf91a88d6a&amp;amp;vid=a2fec0802c3a11f0b5c3b9fd0fc28d87&amp;amp;vids=1&amp;amp;msclkid=N&amp;amp;uach=pv%3D10.0&amp;amp;pi=0&amp;amp;lg=en-US&amp;amp;sw=1280&amp;amp;sh=720&amp;amp;sc=24&amp;amp;nwd=1&amp;amp;tl=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;p=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;r=&amp;amp;lt=569&amp;amp;evt=pageLoad&amp;amp;sv=1&amp;amp;cdb=AQET&amp;amp;rn=206258", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://bat.bing.com/action/0?ti=163002607&amp;amp;Ver=2&amp;amp;mid=ae9bdae8-4702-4299-b1a6-34555b231a31&amp;amp;bo=2&amp;amp;sid=a2fec0b02c3a11f095697fbf91a88d6a&amp;amp;vid=a2fec0802c3a11f0b5c3b9fd0fc28d87&amp;amp;vids=0&amp;amp;msclkid=N&amp;amp;ec=CHEQ&amp;amp;el=Invalid_Users&amp;amp;ev=0&amp;amp;ea=Invalid_Users&amp;amp;en=Y&amp;amp;p=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;sw=1280&amp;amp;sh=720&amp;amp;sc=24&amp;amp;nwd=1&amp;amp;evt=custom&amp;amp;cdb=AQET&amp;amp;rn=414587", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://px.ads.linkedin.com/wa/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;pids&#34;:[64444],&#34;scriptVersion&#34;:213,&#34;time&#34;:1746729081616,&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;url&#34;:&#34;https://crowdstrike.com/en-us/&#34;,&#34;pageTitle&#34;:&#34;CrowdStrike: We Stop Breaches with AI-native Cybersecurity&#34;,&#34;websiteSignalRequestId&#34;:&#34;368743e1-b3ba-3b47-e4d0-caee31d6a196&#34;,&#34;isTranslated&#34;:false,&#34;liFatId&#34;:&#34;&#34;,&#34;liGiant&#34;:&#34;&#34;,&#34;misc&#34;:{&#34;psbState&#34;:-4},&#34;isLinkedInApp&#34;:false,&#34;hem&#34;:null,&#34;signalType&#34;:&#34;CLICK&#34;,&#34;href&#34;:&#34;&#34;,&#34;domAttributes&#34;:{&#34;elementSemanticType&#34;:null,&#34;elementValue&#34;:null,&#34;elementType&#34;:null,&#34;tagName&#34;:&#34;BUTTON&#34;,&#34;backgroundImageSrc&#34;:null,&#34;imageSrc&#34;:null,&#34;imageAlt&#34;:null,&#34;innerText&#34;:&#34;&#34;,&#34;elementTitle&#34;:null,&#34;cursor&#34;:&#34;pointer&#34;},&#34;innerElements&#34;:[{&#34;elementSemanticType&#34;:&#34;IMAGE&#34;,&#34;elementValue&#34;:null,&#34;elementType&#34;:null,&#34;tagName&#34;:&#34;IMG&#34;,&#34;backgroundImageSrc&#34;:null,&#34;imageSrc&#34;:&#34;https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/graphics-and-illustrations/marketplace/addsearch-icon.svg&#34;,&#34;imageAlt&#34;:&#34;Search Icon&#34;,&#34;innerText&#34;:&#34;&#34;,&#34;elementTitle&#34;:null,&#34;cursor&#34;:&#34;pointer&#34;}],&#34;elementCrumbsTree&#34;:[{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:12,&#34;classes&#34;:[&#34;container&#34;,&#34;responsivegrid&#34;,&#34;root&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;id&#34;:&#34;container-5e90248122&#34;,&#34;classes&#34;:[&#34;cmp-container&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;classes&#34;:[&#34;aem-Grid&#34;,&#34;aem-Grid--12&#34;,&#34;aem-Grid--default--12&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;classes&#34;:[&#34;aem-GridColumn&#34;,&#34;aem-GridColumn--default--12&#34;,&#34;experiencefragment&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;id&#34;:&#34;experiencefragment-a599a29fb4&#34;,&#34;classes&#34;:[&#34;cmp-experiencefragment&#34;,&#34;cmp-experiencefragment--header&#34;],&#34;attributes&#34;:{&#34;data-target-location&#34;:&#34;false&#34;}},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;id&#34;:&#34;container-51790f4854&#34;,&#34;classes&#34;:[&#34;cmp-container&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;classes&#34;:[&#34;aem-Grid&#34;,&#34;aem-Grid--12&#34;,&#34;aem-Grid--default--12&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;classes&#34;:[&#34;aem-GridColumn&#34;,&#34;aem-GridColumn--default--12&#34;,&#34;container&#34;,&#34;responsivegrid&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;id&#34;:&#34;container-30992da794&#34;,&#34;classes&#34;:[&#34;cmp-container&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:3,&#34;classes&#34;:[&#34;container&#34;,&#34;container--dotcom-header--sticky&#34;,&#34;responsivegrid&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;id&#34;:&#34;container-ca8a123f47&#34;,&#34;classes&#34;:[&#34;cmp-container&#34;,&#34;container--dotcom-header--sticky__cmp-container--fixed&#34;],&#34;attributes&#34;:{&#34;data-target-location&#34;:&#34;false&#34;}},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;classes&#34;:[&#34;container&#34;,&#34;container--centered&#34;,&#34;container--flex&#34;,&#34;container--horizontal-space-between&#34;,&#34;container--vertical-center&#34;,&#34;responsivegrid&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;id&#34;:&#34;container-eacd549d38&#34;,&#34;classes&#34;:[&#34;cmp-container&#34;],&#34;attributes&#34;:{&#34;data-target-location&#34;:&#34;false&#34;}},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:1,&#34;classes&#34;:[&#34;container&#34;,&#34;container--flex&#34;,&#34;container--vertical-center&#34;,&#34;responsivegrid&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;id&#34;:&#34;container-3be8430239&#34;,&#34;classes&#34;:[&#34;cmp-container&#34;],&#34;attributes&#34;:{&#34;data-target-location&#34;:&#34;false&#34;}},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;classes&#34;:[&#34;embed&#34;]},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;id&#34;:&#34;embed-8a5069951f&#34;,&#34;classes&#34;:[&#34;cmp-embed&#34;],&#34;attributes&#34;:{&#34;data-cmp-data-layer&#34;:&#34;{\&#34;embed-8a5069951f\&#34;:{\&#34;@type\&#34;:\&#34;crowdstrike/components/content/embed/v1/embed\&#34;,\&#34;repo:modifyDate\&#34;:\&#34;2025-05-07T19:34:15Z\&#34;}}&#34;,&#34;data-target-location&#34;:&#34;false&#34;}},{&#34;tagName&#34;:&#34;div&#34;,&#34;nthChild&#34;:0,&#34;id&#34;:&#34;search-header&#34;,&#34;attributes&#34;:{&#34;data-cmp-is&#34;:&#34;add-search&#34;}},{&#34;tagName&#34;:&#34;button&#34;,&#34;nthChild&#34;:0,&#34;classes&#34;:[&#34;search_btn&#34;,&#34;search_btn--redesign&#34;]}],&#34;isFilteredByClient&#34;:false}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://edge.adobedc.net/ee/or2/v1/interact?configId=00798cfe-13d2-4126-bcb1-df59bdd246ce&amp;amp;requestId=e10e27de-cf2c-4743-86b9-b29e80b2a6ea", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;events&#34;:[{&#34;xdm&#34;:{&#34;eventType&#34;:&#34;web.webinteraction.linkClicks&#34;,&#34;web&#34;:{&#34;webPageDetails&#34;:{&#34;URL&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;},&#34;webReferrer&#34;:{&#34;URL&#34;:&#34;&#34;},&#34;webInteraction&#34;:{&#34;URL&#34;:&#34;/en-us/&#34;,&#34;name&#34;:&#34;Search Icon&#34;,&#34;type&#34;:&#34;other&#34;,&#34;region&#34;:&#34;megaSearch&#34;,&#34;linkClicks&#34;:{&#34;value&#34;:1}}},&#34;device&#34;:{&#34;screenHeight&#34;:720,&#34;screenWidth&#34;:1280,&#34;screenOrientation&#34;:&#34;landscape&#34;},&#34;environment&#34;:{&#34;type&#34;:&#34;browser&#34;,&#34;browserDetails&#34;:{&#34;viewportWidth&#34;:1280,&#34;viewportHeight&#34;:720}},&#34;placeContext&#34;:{&#34;localTimezoneOffset&#34;:420,&#34;localTime&#34;:&#34;2025-05-08T11:31:21.618-07:00&#34;},&#34;timestamp&#34;:&#34;2025-05-08T18:31:21.618Z&#34;,&#34;implementationDetails&#34;:{&#34;name&#34;:&#34;https://ns.adobe.com/experience/alloy/reactor&#34;,&#34;version&#34;:&#34;2.26.0&#43;2.29.0&#34;,&#34;environment&#34;:&#34;browser&#34;},&#34;_experience&#34;:{&#34;analytics&#34;:{&#34;event1to100&#34;:{&#34;event41&#34;:{&#34;id&#34;:&#34;Search Icon&#34;,&#34;value&#34;:1}},&#34;customDimensions&#34;:{&#34;eVars&#34;:{&#34;eVar19&#34;:&#34;CrowdStrike-AEM|production|2025-05-05T15:24:36Z|Search Icon Clicked&#34;,&#34;eVar1&#34;:&#34;/en-us/&#34;,&#34;eVar2&#34;:&#34;www.crowdstrike.com&#34;,&#34;eVar3&#34;:&#34;&#34;,&#34;eVar4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;eVar6&#34;:&#34;dir&#34;,&#34;eVar7&#34;:&#34;&#34;,&#34;eVar15&#34;:&#34;28180928486187528650392535887012291588&#34;,&#34;eVar16&#34;:&#34;bg41,c0001,c0002,c0003,c0004&#34;,&#34;eVar17&#34;:&#34;&#34;,&#34;eVar18&#34;:&#34;&#34;,&#34;eVar20&#34;:&#34;/en-us/&#34;,&#34;eVar26&#34;:&#34;in&#34;,&#34;eVar111&#34;:&#34;crowdstrike:homepage&#34;,&#34;eVar112&#34;:&#34;crowdstrike:en-us&#34;,&#34;eVar113&#34;:&#34;&#34;,&#34;eVar114&#34;:&#34;&#34;,&#34;eVar115&#34;:&#34;crowdstrike:en-us&#34;},&#34;props&#34;:{&#34;prop1&#34;:&#34;/en-us/&#34;,&#34;prop2&#34;:&#34;www.crowdstrike.com&#34;,&#34;prop4&#34;:&#34;www.crowdstrike.com/en-us/&#34;,&#34;prop6&#34;:&#34;&#34;,&#34;prop10&#34;:&#34;crowdstrike:homepage&#34;,&#34;prop11&#34;:&#34;crowdstrike:en-us&#34;,&#34;prop12&#34;:&#34;&#34;,&#34;prop13&#34;:&#34;&#34;,&#34;prop14&#34;:&#34;crowdstrike:en-us&#34;}}}}}}],&#34;query&#34;:{&#34;identity&#34;:{&#34;fetch&#34;:[&#34;ECID&#34;,&#34;CORE&#34;]}},&#34;meta&#34;:{&#34;state&#34;:{&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;cookiesEnabled&#34;:true,&#34;entries&#34;:[{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_identity&#34;,&#34;value&#34;:&#34;CiYyODE4MDkyODQ4NjE4NzUyODY1MDM5MjUzNTg4NzAxMjI5MTU4OFISCJ2jqonrMhABGAEqA09SMjAA8AGdo6qJ6zI=&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_cluster&#34;,&#34;value&#34;:&#34;or2&#34;},{&#34;key&#34;:&#34;kndctr_06D71E9261F941560A495CD6_AdobeOrg_consent&#34;,&#34;value&#34;:&#34;general=in&#34;}]}}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.facebook.com/tr/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;id&#34;

992980065451679
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;ev&#34;

SubscribedButtonClick
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;dl&#34;

https://www.crowdstrike.com/en-us/
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;rl&#34;


------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;if&#34;

false
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;ts&#34;

1746729081620
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;cd[buttonFeatures]&#34;

{&#34;classList&#34;:&#34;search_btn search_btn--redesign&#34;,&#34;destination&#34;:&#34;&#34;,&#34;id&#34;:&#34;&#34;,&#34;imageUrl&#34;:&#34;https://www.crowdstrike.com/content/dam/crowdstrike/marketing/en-us/images/graphics-and-illustrations/marketplace/addsearch-icon.svg&#34;,&#34;innerText&#34;:&#34;\n      \n    &#34;,&#34;numChildButtons&#34;:0,&#34;tag&#34;:&#34;button&#34;,&#34;type&#34;:null,&#34;name&#34;:&#34;&#34;,&#34;value&#34;:&#34;&#34;}
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;cd[buttonText]&#34;


      
    
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;cd[formFeatures]&#34;

[]
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;cd[pageFeatures]&#34;

{&#34;title&#34;:&#34;CrowdStrike: We Stop Breaches with AI-native Cybersecurity&#34;}
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;sw&#34;

1280
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;sh&#34;

720
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;v&#34;

2.9.201
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;r&#34;

stable
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;ec&#34;

2
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;o&#34;

4126
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;fbp&#34;

fb.1.1746729079520.75795469452619173
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;ler&#34;

empty
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;it&#34;

1746729079168
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;coo&#34;

false
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;es&#34;

automatic
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;tm&#34;

3
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;exp&#34;

k2
------WebKitFormBoundaryIoyZtne4PWC1gV0m
Content-Disposition: form-data; name=&#34;rqm&#34;

SB
------WebKitFormBoundaryIoyZtne4PWC1gV0m--
`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://c.contentsquare.net/v2/events?uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;amp;sn=1&amp;amp;hd=1746729079&amp;amp;v=15.85.1&amp;amp;pid=29632&amp;amp;pn=1&amp;amp;sr=10&amp;amp;mdh=7994&amp;amp;str=96&amp;amp;di=485&amp;amp;dc=2326&amp;amp;fl=2347&amp;amp;ct=0", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`[{&#34;type&#34;:0,&#34;ts&#34;:7,&#34;x&#34;:1280,&#34;y&#34;:720},{&#34;type&#34;:19,&#34;name&#34;:&#34;FCP&#34;,&#34;val&#34;:516,&#34;ts&#34;:9},{&#34;type&#34;:19,&#34;name&#34;:&#34;LCP&#34;,&#34;val&#34;:1116,&#34;ts&#34;:9},{&#34;type&#34;:19,&#34;name&#34;:&#34;CLS&#34;,&#34;val&#34;:0.7721337735123106,&#34;ts&#34;:10},{&#34;type&#34;:19,&#34;name&#34;:&#34;TTFB&#34;,&#34;val&#34;:152.2000000178814,&#34;ts&#34;:827},{&#34;type&#34;:6,&#34;ts&#34;:868,&#34;x&#34;:1115,&#34;y&#34;:518,&#34;tgt&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;},{&#34;type&#34;:2,&#34;ts&#34;:868,&#34;x&#34;:1115,&#34;y&#34;:518,&#34;xRel&#34;:32768,&#34;yRel&#34;:28603,&#34;tgtHM&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;},{&#34;type&#34;:3,&#34;ts&#34;:869,&#34;x&#34;:1115,&#34;y&#34;:518,&#34;tgt&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;},{&#34;type&#34;:4,&#34;ts&#34;:872,&#34;x&#34;:1115,&#34;y&#34;:518,&#34;tgt&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;},{&#34;type&#34;:5,&#34;ts&#34;:874,&#34;x&#34;:1115,&#34;y&#34;:518,&#34;tgt&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;,&#34;xRel&#34;:32768,&#34;yRel&#34;:28603},{&#34;type&#34;:19,&#34;name&#34;:&#34;INP&#34;,&#34;val&#34;:48,&#34;ts&#34;:932},{&#34;type&#34;:19,&#34;name&#34;:&#34;FID&#34;,&#34;val&#34;:1.2999999821186066,&#34;ts&#34;:932},{&#34;type&#34;:3,&#34;ts&#34;:953,&#34;x&#34;:1115,&#34;y&#34;:518,&#34;tgt&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;},{&#34;type&#34;:4,&#34;ts&#34;:953,&#34;x&#34;:1115,&#34;y&#34;:518,&#34;tgt&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;},{&#34;type&#34;:5,&#34;ts&#34;:953,&#34;x&#34;:1115,&#34;y&#34;:518,&#34;tgt&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;,&#34;xRel&#34;:32768,&#34;yRel&#34;:28603},{&#34;type&#34;:2,&#34;ts&#34;:1269,&#34;x&#34;:1115,&#34;y&#34;:518,&#34;xRel&#34;:32768,&#34;yRel&#34;:28603,&#34;tgtHM&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;},{&#34;type&#34;:7,&#34;ts&#34;:1298,&#34;x&#34;:1115,&#34;y&#34;:526,&#34;tgt&#34;:&#34;div#onetrust-close-btn-container&gt;button:eq(0)&#34;},{&#34;type&#34;:1,&#34;ts&#34;:1585,&#34;x&#34;:0,&#34;y&#34;:93,&#34;d&#34;:304},{&#34;type&#34;:6,&#34;ts&#34;:1900,&#34;x&#34;:917,&#34;y&#34;:127,&#34;tgt&#34;:&#34;div#search-header&gt;button:eq(0)&gt;img:eq(0)&#34;},{&#34;type&#34;:2,&#34;ts&#34;:1900,&#34;x&#34;:917,&#34;y&#34;:127,&#34;xRel&#34;:32460,&#34;yRel&#34;:32768,&#34;tgtHM&#34;:&#34;div#search-header&gt;button:eq(0)&gt;img:eq(0)&#34;},{&#34;type&#34;:3,&#34;ts&#34;:1900,&#34;x&#34;:917,&#34;y&#34;:127,&#34;tgt&#34;:&#34;div#search-header&gt;button:eq(0)&gt;img:eq(0)&#34;},{&#34;type&#34;:4,&#34;ts&#34;:1902,&#34;x&#34;:917,&#34;y&#34;:127,&#34;tgt&#34;:&#34;div#search-header&gt;button:eq(0)&gt;img:eq(0)&#34;},{&#34;type&#34;:5,&#34;ts&#34;:1905,&#34;x&#34;:917,&#34;y&#34;:127,&#34;tgt&#34;:&#34;div#search-header&gt;button:eq(0)&gt;img:eq(0)&#34;,&#34;xRel&#34;:32460,&#34;yRel&#34;:32768}]`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://api.addsearch.com/v1/search/7737a29b854de71521b1cd72c4118cfc?term=_addsearch_0.18653986968752023&amp;amp;fuzzy=auto&amp;amp;collectAnalytics=false&amp;amp;page=1&amp;amp;limit=10&amp;amp;sort=relevance&amp;amp;order=desc&amp;amp;filter=%7B%7D", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "application/json, text/plain, */*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/snowflake/analyticsdata?timestamp=1746729081643", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;event-name&#34;:&#34;close cart&#34;,&#34;tag&#34;:&#34;Button&#34;,&#34;tag-location&#34;:&#34;Cart Icon&#34;,&#34;landingpage-url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;window-inner-width&#34;:1280,&#34;window-inner-height&#34;:720,&#34;body-client-width&#34;:1280,&#34;body-client-height&#34;:720,&#34;screen-width&#34;:1280,&#34;screen-height&#34;:720,&#34;screen-width-height&#34;:&#34;1280x720&#34;,&#34;device-pixel-ratio&#34;:1,&#34;browser-name&#34;:&#34;Chrome&#34;,&#34;browser-version&#34;:&#34;136.0.7103.25&#34;,&#34;os-name&#34;:&#34;Windows&#34;,&#34;os-version&#34;:&#34;NT 10.0&#34;,&#34;engine-name&#34;:&#34;Blink&#34;,&#34;engine-version&#34;:&#34;&#34;,&#34;device-type&#34;:&#34;desktop&#34;,&#34;session_id&#34;:&#34;8ffaa7ae-9b97-463d-b437-ce63a0483042&#34;,&#34;productName&#34;:&#34;&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=58375e44-e5e0-4f70-bfd4-94cff2da01b0&amp;amp;batch_time=1746729081680", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080775,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;8010c39d-9714-4bc9-9ce0-73d3a346c0d4&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:229000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://js.zi-scripts.com/unified/v1/master/getSubscriptions&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080793,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;2b080a39-99ad-43d4-9106-a1e6085a120b&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:211000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://grsm.io/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23?get_pscd=true&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080935,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;df22892e-4b23-4b0b-8582-6872273adab1&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:111000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://partnerlinks.io/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080935,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;4d9a2992-7c7c-4cb8-b910-2d6b686b1707&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:129000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://get.crowdstrike.com/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080880,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b3ed4d76-32ed-45a0-818e-6727bf15b86b&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:202000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://siteintercept.qualtrics.com/WRSiteInterceptEngine/Targeting.php?Q_ZoneID=ZN_dnXlCJ0uLH6d1ZQ&amp;Q_CLIENTVERSION=2.28.0&amp;Q_CLIENTTYPE=webAdobeLaunch&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081004,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5d1c3aee-5f12-4707-9967-d6bf8c66e841&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:79000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://partnerlinks.io/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081004,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1640902d-9502-4eba-8738-cd53cb7289c3&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:93000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://get.crowdstrike.com/pr/grc/pk_BntOCdi156BS0syKkTE4XezLhvMIjQ23&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080775,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;bc05af96-235a-4f98-8a91-8cc32c5042cd&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:323000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://js.zi-scripts.com/unified/v1/master/getSubscriptions&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081069,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b7bd04ee-1fa7-4a76-a15a-3f4a7f11d39b&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:97000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://obs.fishrobotflower.com/mon&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081074,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;16159553-1fcb-4c6e-b2d2-011d70d0aad3&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:104000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://obs.fishrobotflower.com/mon&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081428,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;414a68f7-b4cb-4808-9a6b-1b757b3048a0&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://bat.bing.com/action/0?ti=163002607&amp;Ver=2&amp;mid=ae9bdae8-4702-4299-b1a6-34555b231a31&amp;bo=1&amp;sid=a2fec0b02c3a11f095697fbf91a88d6a&amp;vid=a2fec0802c3a11f0b5c3b9fd0fc28d87&amp;vids=1&amp;msclkid=N&amp;uach=pv%3D10.0&amp;pi=0&amp;lg=en-US&amp;sw=1280&amp;sh=720&amp;sc=24&amp;nwd=1&amp;tl=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;p=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;r=&amp;lt=569&amp;evt=pageLoad&amp;sv=1&amp;cdb=AQET&amp;rn=206258&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:37000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081428,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;eb040d13-8dee-4d62-a318-ca2ca51beb97&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://bat.bing.com/action/0?ti=163002607&amp;Ver=2&amp;mid=ae9bdae8-4702-4299-b1a6-34555b231a31&amp;bo=2&amp;sid=a2fec0b02c3a11f095697fbf91a88d6a&amp;vid=a2fec0802c3a11f0b5c3b9fd0fc28d87&amp;vids=0&amp;msclkid=N&amp;ec=CHEQ&amp;el=Invalid_Users&amp;ev=0&amp;ea=Invalid_Users&amp;en=Y&amp;p=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;sw=1280&amp;sh=720&amp;sc=24&amp;nwd=1&amp;evt=custom&amp;cdb=AQET&amp;rn=414587&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:49500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080918,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;3ded5084-500a-4a1d-8994-270220b96db9&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:727000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://ws.zoominfo.com/pixel/61b22df2e97826001a6d4b6e/?iszitag=true&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081621,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;311f426b-3c12-48c4-b334-c3030784323c&#34;,&#34;type&#34;:&#34;beacon&#34;,&#34;url&#34;:&#34;https://www.facebook.com/tr/&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:22700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081418,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;f03a7502-8a32-4c87-968e-90ab33e1ca75&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/ipv?_biz_r=&amp;_biz_h=-1719904874&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079103&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=0&amp;a=crowdstrike.com&amp;rnd=480761&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:252900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081617,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;914f3b44-8dc1-402c-a9b7-fa36c69ad9f6&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:54000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:204,&#34;url&#34;:&#34;https://px.ads.linkedin.com/wa/&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://api.addsearch.com/v1/search/7737a29b854de71521b1cd72c4118cfc?term=blash&amp;amp;fuzzy=auto&amp;amp;collectAnalytics=false&amp;amp;page=1&amp;amp;limit=10&amp;amp;sort=relevance&amp;amp;order=desc&amp;amp;filter=%7B%7D", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept", "application/json, text/plain, */*")
	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://ws.zoominfo.com/pixel/collect", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;eventId&#34;:&#34;e77d3f60-eca4-4684-9c50-6f4b6804d79a&#34;,&#34;websiteId&#34;:&#34;61b22df2e97826001a6d4b6e&#34;,&#34;companyId&#34;:&#34;206553737&#34;,&#34;sessionId&#34;:&#34;8733815197353b9f320ca675cb207cc32c56bae4c3f0f504d4931c9b00b77bfa&#34;,&#34;visitorId&#34;:&#34;60cb6f73ec53efe523361746729080&#34;,&#34;eventCreatedAt&#34;:&#34;2025-05-08T18:31:21.495Z&#34;,&#34;secs&#34;:0,&#34;v&#34;:7}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://event.api.drift.com/track", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;orgId&#34;:113673,&#34;inboxId&#34;:113963,&#34;userId&#34;:null,&#34;attributes&#34;:{&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;,&#34;title&#34;:&#34;CrowdStrike: We Stop Breaches with AI-native Cybersecurity&#34;,&#34;pageName&#34;:&#34;&#34;},&#34;event&#34;:&#34;Page&#34;,&#34;embedId&#34;:&#34;9d4udx6ceimp&#34;,&#34;context&#34;:{&#34;userAgent&#34;:&#34;Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.25 Safari/537.36&#34;},&#34;anonymousId&#34;:&#34;3a55016c-317c-4ceb-935c-33bce08ec640&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://targeting.api.drift.com/targeting/evaluate_with_log", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;conditionGroups&#34;:[{&#34;status&#34;:&#34;EVALUATED&#34;,&#34;matched&#34;:true,&#34;conditionSets&#34;:[{&#34;targetingConditions&#34;:[{&#34;evaluatedCondition&#34;:{&#34;field&#34;:&#34;urlPath&#34;,&#34;operator&#34;:&#34;includesAnyOf&#34;,&#34;value&#34;:[&#34;/platform/&#34;,&#34;/platform/next-gen-siem/&#34;,&#34;/en-us/&#34;,&#34;/platform/next-gen-siem/falcon-fusion/&#34;,&#34;/platform/next-gen-siem/falcon-foundry/&#34;,&#34;/platform/next-gen-siem/falcon-logscale/&#34;,&#34;/platform/next-gen-siem/falcon-search-retention/&#34;,&#34;/services/falcon-complete-next-gen-mdr/&#34;,&#34;/en-us/services/&#34;,&#34;/en-us/resources/data-sheets/crowdstrike-falcon-next-gen-siem/&#34;,&#34;/blog/category/observability-and-log-management/&#34;,&#34;/blog/augment-or-replace-siem-with-crowdstrike-falcon/&#34;,&#34;/blog/top-5-siem-use-cases-logscale-solves/&#34;,&#34;/blog/building-modern-soc-with-next-gen-siem/&#34;,&#34;/blog/montage-health-consolidates-cybersecurity-strategy-with-crowdstrike/&#34;,&#34;/blog/four-falcon-logscale-ng-siem-updates/&#34;,&#34;/blog/accelerate-response-with-falcon-fusion/&#34;,&#34;/blog/falcon-logscale-chrome-enterprise-security-telemetry-integration/&#34;,&#34;/blog/move-from-legacy-to-next-gen-siem-falcon-logscale/&#34;,&#34;/en-us/blog/leveraging-crowdstrike-falcon-against-attacks-targeting-okta/&#34;,&#34;/en-us/blog/leveraging-crowdstrike-falcon-against-attacks-targeting-okta/&#34;,&#34;/en-us/resources/white-papers/future-proof-your-soc-migration-guide/&#34;,&#34;/en-us/resources/reports/forrester-wave-cybersecurity-incident-response-services-2024/&#34;,&#34;/en-us/blog/category.from-the-front-lines/&#34;,&#34;/en-us/resources/crowdcasts/reactive-to-proactive-cybersecurity-transformation/&#34;,&#34;/en-us/resources/white-papers/5-services-retainer-scenarios/&#34;,&#34;/en-us/resources/white-papers/optimize-your-profile-for-cyber-insurance/&#34;,&#34;/en-us/resources/customer-stories/tabcorp-partners-with-crowdstrike/&#34;,&#34;/en-us/resources/customer-stories/jemena/&#34;,&#34;/en-us/resources/customer-stories/nij-smellinghe-video/&#34;,&#34;/en-us/resources/customer-stories/uk-university-crowdstrike-mdr/&#34;,&#34;/en-us/resources/customer-stories/tdk-electronics/&#34;,&#34;/en-us/resources/customer-stories/telus-health-video/&#34;,&#34;/en-us/resources/crowdcasts/harness-total-enterprise-visibility-and-protection-with-next-gen-mdr/&#34;,&#34;/en-us/resources/guides/managed-security-readiness-guide/&#34;,&#34;/en-us/resources/reports/2024-frost-sullivan-mdr-company-of-the-year-mdr/&#34;,&#34;/en-us/resources/reports/2024-gartner-peer-insights-voc-mdr-report/&#34;,&#34;/en-us/resources/reports/forrester-wave-mdr-services-q1-2025/&#34;,&#34;/en-us/resources/customer-stories/rate-companies/&#34;,&#34;/en-us/resources/customer-stories/aerospace-tech-company/&#34;,&#34;/en-us/resources/customer-stories/&#34;,&#34;/en-us/resources/white-papers/state-of-the-siem-market/&#34;,&#34;/en-us/resources/guides/next-gen-siem-rfp-checklist/&#34;,&#34;/en-us/blog/unlocking-soc-superpowers/&#34;,&#34;/en-us/blog/unlocking-soc-superpowers/&#34;,&#34;/en-us/blog/need-for-speed-soc/&#34;,&#34;/en-us/blog/unlock-advanced-security-automation-for-next-gen-siem/&#34;,&#34;/blog/&#34;,&#34;/cybersecurity-101/observability/siem-vs-log-management/&#34;,&#34;/cybersecurity-101/next-gen-siem/&#34;,&#34;/cybersecurity-101/security-information-and-event-management-siem/&#34;,&#34;/cybersecurity-101/soc-automation/&#34;],&#34;providerName&#34;:null},&#34;matched&#34;:true,&#34;status&#34;:&#34;EVALUATED&#34;},{&#34;evaluatedCondition&#34;:{&#34;field&#34;:&#34;urlPath&#34;,&#34;operator&#34;:&#34;notIncludesAnyOf&#34;,&#34;value&#34;:[&#34;/products/next-gen-siem&#34;,&#34;/careers&#34;,&#34;/explore&#34;,&#34;/future&#34;,&#34;/future_roxe02gowrispo5r99&#34;,&#34;/ar-sa&#34;,&#34;/nab&#34;,&#34;/pt-br&#34;,&#34;/de-de&#34;,&#34;/es-es&#34;,&#34;/fr-fr&#34;,&#34;/es-latam&#34;,&#34;/ja-jp&#34;,&#34;/zh-tw&#34;,&#34;/ko-kr&#34;,&#34;/it-it&#34;,&#34;/contact-us-kr&#34;,&#34;/contact-us-tw&#34;,&#34;/contact-us-ar&#34;,&#34;/contato&#34;,&#34;/contact-us-it&#34;,&#34;/kontakt&#34;,&#34;/contact-us-es&#34;,&#34;/contact-us&#34;,&#34;/request-information&#34;,&#34;/explore/&#34;,&#34;/crowdstrike-financial-services/&#34;,&#34;/1password&#34;,&#34;/platform/saas-security-posture-management/&#34;],&#34;providerName&#34;:null},&#34;matched&#34;:true,&#34;status&#34;:&#34;EVALUATED&#34;},{&#34;evaluatedCondition&#34;:{&#34;field&#34;:&#34;device&#34;,&#34;operator&#34;:&#34;notIsAnyOf&#34;,&#34;value&#34;:[&#34;mobile&#34;,&#34;tablet&#34;],&#34;providerName&#34;:null},&#34;matched&#34;:true,&#34;actualValue&#34;:&#34;desktop&#34;,&#34;status&#34;:&#34;EVALUATED&#34;},{&#34;evaluatedCondition&#34;:{&#34;field&#34;:&#34;clearbit&#34;,&#34;operator&#34;:&#34;clearbit&#34;,&#34;value&#34;:&#34;AGGrlcI7bCOp5zrjJuFshdhzua7SBzI37Vm12cPvHJnHLch5MjczgaEj1IXL6kIHCqfJFptneRhj9szewqi01y1hnzZGN8LsfYigV5lr0BfkLAWxuYCTRzNL-UQjTq7wzSgb5xtIwXS4I8b5-Cd3wwZKRPUT8cNG1v40zWXZy_3UMKIC_ovqLk6Cue4h0ZMC230Z1jG45PmV2mwBCkU4p8WTXX_7unz_K1LweSUFz2h4a7E7VpjRfcM&#34;,&#34;providerName&#34;:null},&#34;matched&#34;:true,&#34;status&#34;:&#34;EVALUATED&#34;},{&#34;evaluatedCondition&#34;:{&#34;field&#34;:&#34;clearbit&#34;,&#34;operator&#34;:&#34;clearbit&#34;,&#34;value&#34;:&#34;AGGrlcJFlWPbZP4HVRe4ebj1jMCKyxGoQnkjvhzf--HO_d8WBf3KLH3i8RhLwrJW6KAkizzRpP0EJPBCPSZUCVGFUnTxI9O3zd5m4ZjCuAgR6O5ttsraHamq2UI0xgeTHA4NV2aBG8UPXY3ceCq8kKb8qEnFMtldAM6YzZvSEM28t8iIqTvwIz7P3CSVRnw-dPcwiztKQz7O5hHCUtxcxyKg2BIG2UPzPA&#34;,&#34;providerName&#34;:null},&#34;matched&#34;:true,&#34;status&#34;:&#34;EVALUATED&#34;}],&#34;status&#34;:&#34;EVALUATED&#34;}],&#34;key&#34;:2924894,&#34;priority&#34;:1}],&#34;targetingContext&#34;:{&#34;endUserId&#34;:25104570104,&#34;orgId&#34;:113673},&#34;evaluationOptions&#34;:{&#34;ignoreUnknowns&#34;:false}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://obs.fishrobotflower.com/mon", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`e=37dfbd8ee84e00126ee8c037e347829d9225c24f567d43d6da1908be6245cad7bd70a976710ce60ed89373bfe70e9c20c1e53e8d5a138c6c2717071a10acf9f29f671a82d589562b6d1ef878215780338a33c103300376c5065359600d5ccfbf394f77be26bb25cb43e2913df05365ac0b297d1bdc57e446f492d7da3dbb280ef67ccaa8073f8d0e3717224793845731f360b3f493a0180dec1edae97dfa2bc8169b1adc597cff3200e714561c4b92177af998ffe4198b6dec06c213f85e162ae7d133722b325f817c99ec59b058609fc6e359143e3dd385293e88864c06513c157a77bb9e70392652b48d1c2ad7f4ec3ee3b8192d4079b4a7a6928677a0dadf5ae8489e5d2e019cbecbf7af2b95dfe57594351ccdeb8b795904fd7367a095166dbdd0aa73d902fd12b0801d91499d9978953b67919d27dc6796d3bbe193fdbd4c38fc28b0bdfd354371fe8f719aa61af7010642dd4245c2979684c0ec8825ce3bc68ebe817fa27d763bc97f27db838003679ea875ac41bfc003d667a87e5346588c6eee5f8b89ee7ef8df78f02da7c19cdb27d60a0a2dea5d540b28073ed7ab67b355c4cb9764fdde25b3d02dd0c2b1be7c85b7f409746c89a91004d7b41d5107f7fae939fa2890f802fdfc3ce52dd36dcc068a50aefc2ecd0823e2ac90633f3d6402340a5b13ace68e65d205be5752517faac71386cd464887eccdb0ae6ddcece3f880ffadf84b4835fe4586a85e6b844a601fd16a59746ef07aa6771ec1f9dd71c83cc52fe0a82cc8a40e99721322210760129cc0e4db207a559a184689d2d9bbb0c24f049eb0be7f8ad2757adbce0f3a68fd8dc45020b23b98599e1cff22981ad8d40efbb337fe868bd49e5b7feee3337f9f71a68aacb6119d6d280b02956a5f4db6a25b93f112ca9887d552cd839b2e083745d321b951ef0ce6936384335d20463e7b78e3538e863a4bcf1cc7c2f03da8cb0be60d76e678072c8df4c3491fb725e78402db1930074eba3c2e8bafd061d52c49e4ff15ebbd9d88e11aa90924605fe06b3716d9c9584eb02ee6e0c02a8bd5eb0c547ca5633e507e867459e6c8e66c40cd260bbbd6c353d5c6f88f0f30b05bd61b781488173175d4cc1615da92b2aabde142a563f930b77096d65d9afff22a2bfdd215808cb4b1c4f6cadf741fddf42f54872dd257c81cf3ee28a8cad1c731efc4c7f2fb5a69cba2b5af51cb5594404921413d88823335dc60f64f8ac965ef0522e2714c9766e69c7575b97521525796f3d12f87c74c603d1e6ac2d583ffce74db&amp;cri=jfECv66HI9&amp;sf=0&amp;dc=&amp;cp=custom&amp;gtm=WyJPbmVUcnVzdExvYWRlZCIsIk9wdGFub25Mb2FkZWQiLCJPbmVUcnVzdEdyb3Vwc1VwZGF0ZWQiLCJjb252ZXJzaW9uIiwicGFnZV92aWV3Il0%3D&amp;gac=462250241.1746729079&amp;mm=3%7C1115%2C518%2C0%2C0%2C1115%2C518%2C518%2C1%2C1115%2C518%7C917%2C34%2C-198%2C-484%2C917%2C127%2C1549%2C1%2C917%2C34&amp;md=3%7C1115%2C518%2C1115%2C518%2C520%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C917%2C34%2C917%2C127%2C1551%2C1%2CYnV0dG9uLDQ0LDY4LDg5NSwwLDAsMA%3D%3D%2C%2C917%2C34&amp;mu=3%7C1115%2C518%2C1115%2C518%2C521%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C917%2C34%2C917%2C127%2C1552%2C1%2CYnV0dG9uLDQ0LDY4LDg5NSwwLDAsMA%3D%3D%2C%2C917%2C34&amp;cl=3%7C1115%2C518%2C1115%2C518%2C528%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C917%2C34%2C917%2C127%2C1576%2C1%2CYnV0dG9uLDQ0LDY4LDg5NSwwLDAsMA%3D%3D%2C%2C917%2C34&amp;tb=1&amp;gi=9%2C2448%7C12%2C2449%7C10%2C2450%7C14%2C2452%7C11%2C2452%7C13%2C2452%7C21%2C2461%7C18%2C2862&amp;ws=1280x720&amp;wos=1280x720&amp;ver=13&amp;fi=517&amp;ti=2002&amp;mo=0&amp;pn=3933&amp;spn=1931&amp;fp=516&amp;snt=1&amp;sc=10%7C0%2C2%2C931%2C1%7C0%2C93%2C1080%2C1`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.addsearch.com/v5/assets/search-ui-filetype-icon-pdf.svg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/haas_grot_disp/HaasGrotDisp-75Bold.woff2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/ab4b7b6f9a506e731d820031d0d59176-20250404.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/8286945929327c4f052d4cbfa8b74fed-20250403.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/a967dab279462cd1cebe42200b4f2f33-20250228.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/2a767faa4850d16352fcab54793a9929-20250227.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/6ef1b26d9f40cf23fafeb151b569249b-20250501.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/2a767faa4850d16352fcab54793a9929-20250302.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/b47efe1afc5f0e84ad9f182e2e47eab7-20250228.jpg", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://ws.zoominfo.com/pixel/collect", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;eventId&#34;:&#34;984b9f3d-9ff0-4194-94bf-ef1158d6a979&#34;,&#34;websiteId&#34;:&#34;61b22df2e97826001a6d4b6e&#34;,&#34;companyId&#34;:&#34;206553737&#34;,&#34;sessionId&#34;:&#34;8733815197353b9f320ca675cb207cc32c56bae4c3f0f504d4931c9b00b77bfa&#34;,&#34;visitorId&#34;:&#34;0381becc0335a7bf5d5e1746729081&#34;,&#34;eventCreatedAt&#34;:&#34;2025-05-08T18:31:21.777Z&#34;,&#34;secs&#34;:0,&#34;v&#34;:7}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://targeting.api.drift.com/impressions/widget", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;orgId&#34;:113673,&#34;endUserId&#34;:25104570104,&#34;sessionId&#34;:&#34;bb076d85-176d-40bf-8edd-e4d09c550682&#34;,&#34;sentAt&#34;:1746729082176,&#34;impressionSource&#34;:{&#34;source&#34;:&#34;ACTIVE_PLAYBOOK&#34;,&#34;playbookId&#34;:2828029,&#34;instanceId&#34;:&#34;5e2d5498-aa99-4851-a408-bdc0dce124ee&#34;,&#34;virtual&#34;:true,&#34;widgetVersion&#34;:2,&#34;sendId&#34;:&#34;ALWAYS:25104570104:2828029:5e2d5498-aa99-4851-a408-bdc0dce124ee&#34;,&#34;timeToImpression&#34;:1177,&#34;delayed&#34;:false,&#34;playbookVersion&#34;:1746207802603,&#34;isFromConversationalLandingPage&#34;:false,&#34;loadingStrategy&#34;:&#34;DEFAULT&#34;,&#34;iframeMode&#34;:false,&#34;conversationId&#34;:null},&#34;context&#34;:{&#34;url&#34;:&#34;crowdstrike.com/en-us&#34;,&#34;userAgent&#34;:&#34;Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.25 Safari/537.36&#34;,&#34;country&#34;:&#34;US&#34;}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://c.contentsquare.net/dvar?v=15.85.1&amp;amp;pid=29632&amp;amp;pn=1&amp;amp;sn=1&amp;amp;uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;amp;dv=H4sIAAAAAAAAA6tWcvaIdwmKD8hJrEzKz89WcMssSk1RslJyTswtSMxMz1PwdLFSMLI0MrGwNFGqBQDRyonVLwAAAA%3D%3D&amp;amp;ct=2&amp;amp;r=868288", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/60.be788cda.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=3c320038-1377-40c8-b069-82ca89185cb5&amp;amp;batch_time=1746729082189", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081418,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;ce6e66be-6726-4fd3-a6df-c62874354c17&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizibly.com/u?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079104&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;a=crowdstrike.com&amp;rnd=110587&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:254100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081418,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;3fadb583-01cd-49c5-a70c-b09234d246c5&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/u?mapType=mkto&amp;mapValue=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729080105&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=1&amp;a=crowdstrike.com&amp;rnd=598845&amp;cdn_o=a&amp;_biz_z=1746729080105&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:255200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081418,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;0541cc0a-331e-4123-92fb-b1b923e6e91d&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/u?mapType=ecid&amp;mapValue=06D71E9261F941560A495CD6%40AdobeOrg_28180928486187528650392535887012291588&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729080106&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=2&amp;a=crowdstrike.com&amp;rnd=228061&amp;cdn_o=a&amp;_biz_z=1746729080208&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:256100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081632,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b68228bd-7fc5-49de-b6f1-ed35d268e3a4&#34;,&#34;type&#34;:&#34;beacon&#34;,&#34;url&#34;:&#34;https://c.contentsquare.net/v2/events?uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;sn=1&amp;hd=1746729079&amp;v=15.85.1&amp;pid=29632&amp;pn=1&amp;sr=10&amp;mdh=7994&amp;str=96&amp;di=485&amp;dc=2326&amp;fl=2347&amp;ct=0&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:90300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:0,&#34;encoded_body_size&#34;:0,&#34;decoded_body_size&#34;:0,&#34;transfer_size&#34;:300,&#34;download&#34;:{&#34;duration&#34;:100000,&#34;start&#34;:90200000},&#34;first_byte&#34;:{&#34;duration&#34;:89900000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081005,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9fa7d9da-496a-45b5-9b83-ca299b608824&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:898000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://ws.zoominfo.com/pixel/61b22df2e97826001a6d4b6e/?iszitag=true&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081637,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;69721063-cda9-4d54-868b-8b0cd155fff8&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:317000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://api.addsearch.com/v1/search/7737a29b854de71521b1cd72c4118cfc?term=_addsearch_0.18653986968752023&amp;fuzzy=auto&amp;collectAnalytics=false&amp;page=1&amp;limit=10&amp;sort=relevance&amp;order=desc&amp;filter=%7B%7D&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081643,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;f4ed7143-37ec-4686-80cb-173ef0633daa&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:360800000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/snowflake/analyticsdata?timestamp=1746729081643&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:153,&#34;encoded_body_size&#34;:127,&#34;decoded_body_size&#34;:153,&#34;transfer_size&#34;:427,&#34;download&#34;:{&#34;duration&#34;:300000,&#34;start&#34;:360500000},&#34;first_byte&#34;:{&#34;duration&#34;:359900000,&#34;start&#34;:600000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081880,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;ed84909f-16c1-448a-8c6f-73c89355bf5a&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:201000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://api.addsearch.com/v1/search/7737a29b854de71521b1cd72c4118cfc?term=blash&amp;fuzzy=auto&amp;collectAnalytics=false&amp;page=1&amp;limit=10&amp;sort=relevance&amp;order=desc&amp;filter=%7B%7D&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082086,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;75c3e440-7737-4fd4-9ca8-53cc8c8a0c16&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.addsearch.com/v5/assets/search-ui-filetype-icon-pdf.svg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:60500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082087,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;154c9245-70db-40c7-8a72-cb256ed9531e&#34;,&#34;type&#34;:&#34;font&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/etc.clientlibs/crowdstrike/clientlibs/crowdstrike-fonts/resources/fonts/haas_grot_disp/HaasGrotDisp-75Bold.woff2&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:60000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:33620,&#34;encoded_body_size&#34;:33648,&#34;decoded_body_size&#34;:33620,&#34;transfer_size&#34;:33948,&#34;download&#34;:{&#34;duration&#34;:2100000,&#34;start&#34;:57900000},&#34;first_byte&#34;:{&#34;duration&#34;:54100000,&#34;start&#34;:3800000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081098,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;774e8d14-1a63-4244-bc77-d7f50d129845&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:1053000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://ws.zoominfo.com/pixel/61b22df2e97826001a6d4b6e/?iszitag=true&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082068,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;cd8fc712-87ec-4670-8a83-a1705f3c68ff&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:97000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://obs.fishrobotflower.com/mon&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081904,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c4b48285-8c01-40fe-9029-81fb495b7901&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:262000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://ws.zoominfo.com/pixel/collect&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082089,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;fc0fcd3f-4921-4e27-b0f7-8b6e9cb58fd8&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/2a767faa4850d16352fcab54793a9929-20250302.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:94900000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/60.be788cda.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://driftt.imgix.net/https%3A%2F%2Fs3.us-east-1.amazonaws.com%2Fbot-avatars-prod%2F91%2Ff4548d154d5d5986648a18fc00b02959gt5yv43iw4fz?fit=max&amp;amp;fm=png&amp;amp;h=200&amp;amp;w=200&amp;amp;s=f11ea4a278b492dd53b5be4278ea15f5", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/css/37.102700a2.chunk.css", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/core/assets/js/37.cd6923e6.chunk.js", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/deploy/assets/static/fonts/memvYaGs126MiZpBA-UvWbX2vVnXBbObj2OVTS-mu0SC55I.woff2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://js.driftt.com/deploy/assets/static/fonts/memvYaGs126MiZpBA-UvWbX2vVnXBbObj2OVTS-mu0SC55I.woff2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/ipv?_biz_r=&amp;amp;_biz_h=-1719904874&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729079103&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=0&amp;amp;a=crowdstrike.com&amp;amp;rnd=480761&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729079104", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizibly.com/u?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729079104&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;a=crowdstrike.com&amp;amp;rnd=110587&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729079104", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/u?mapType=mkto&amp;amp;mapValue=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729080105&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=1&amp;amp;a=crowdstrike.com&amp;amp;rnd=598845&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729080105", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/u?mapType=ecid&amp;amp;mapValue=06D71E9261F941560A495CD6%40AdobeOrg_28180928486187528650392535887012291588&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729080106&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=2&amp;amp;a=crowdstrike.com&amp;amp;rnd=228061&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729080208", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/ipv?_biz_r=&amp;amp;_biz_h=-1719904874&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729079103&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=0&amp;amp;a=crowdstrike.com&amp;amp;rnd=480761&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729079104", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizibly.com/u?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729079104&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;a=crowdstrike.com&amp;amp;rnd=110587&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729079104", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/u?mapType=mkto&amp;amp;mapValue=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729080105&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=1&amp;amp;a=crowdstrike.com&amp;amp;rnd=598845&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729080105", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn.bizible.com/u?mapType=ecid&amp;amp;mapValue=06D71E9261F941560A495CD6%40AdobeOrg_28180928486187528650392535887012291588&amp;amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;_biz_t=1746729080106&amp;amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_biz_n=2&amp;amp;a=crowdstrike.com&amp;amp;rnd=228061&amp;amp;cdn_o=a&amp;amp;_biz_z=1746729080208", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn77.api.userway.org/api/img-dscr/v2/dyvvHf6oG0/2376540/V4sorvbtJTNsh59B/alts.json?dto=%7B%22sorted%22%3A%5B%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fblack-primary-crowdstrike-logo-1-addedPadding-3%3Fts%3D1746646455211%26dpr%3Doff%22%2C%22alt%22%3A%22CrowdStrike%20Logo%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FCloud-Security-2%22%2C%22alt%22%3A%22Cloud%20Security%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FData-Protection-3%22%2C%22alt%22%3A%22Data%20Protection%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fempty-cart-image%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FEndpoint-Security-2%22%2C%22alt%22%3A%22Endpoint%20Security%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fexperience-breach-Icon%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FExposure-Management-2%22%2C%22alt%22%3A%22Exposure%20Management%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FGenerative-AI%22%2C%22alt%22%3A%22Generative%20AI%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FIdentity-Protection-1%22%2C%22alt%22%3A%22Identity%20Protection%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FIT-Automation%22%2C%22alt%22%3A%22IT%20Automation%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Flogin-icon%22%2C%22alt%22%3A%22Login%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FNext-Gen-SIEM%22%2C%22alt%22%3A%22Next-Gen%20SIEM%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fphone-icon%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FPlatform-4-16-25-darkmode-HP%3Fts%3D1746621981798%26dpr%3Doff%22%2C%22alt%22%3A%22platform%20graphic%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fprivacyoptions-icon-footer%22%2C%22alt%22%3A%22Your%20Privacy%20Choices%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FSaaS-Security%22%2C%22alt%22%3A%22SaaS%20Security%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fshopping-cart-empty%22%2C%22alt%22%3A%22cart%20icon%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FThreat-Intelligence-Hunting%22%2C%22alt%22%3A%22Threat%20Intelligence%20%26%20Hunting%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FWorkflow-Automation%22%2C%22alt%22%3A%22Workflow%20Automation%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fxiot-security%22%2C%22alt%22%3A%22XIoT%20Security%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2F25-SRV-007_Forrester%2520Wave%2520MDR%25203153x3860%3Fts%3D1746621981959%26dpr%3Doff%22%2C%22alt%22%3A%22Forrester%20MDR%202025%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fcharlotte-hp-hero-1%3F%26hei%3D769%26wid%3D1600%26qlt%3D90%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fcloud-lock-generic-img%3Fts%3D1746621982750%26dpr%3Doff%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Ffrost-radar-2024-2%3Fts%3D1746621982168%26dpr%3Doff%22%2C%22alt%22%3A%22Frost%20CNAPP%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fgartner-mq-leader-report%3Fts%3D1746621982064%26dpr%3Doff%22%2C%22alt%22%3A%22Gigaom%20Ransomware%20Protection%20Chart%202024%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fidentity-security-services-new%3Fts%3D1746621982825%26dpr%3Doff%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fsoc-survival-guide-2.bak%3Fts%3D1746621982899%26dpr%3Doff%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Ftry-free-footer%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fcdn.cookielaw.org%2Flogos%2Fc109dae9-46f3-4e91-a59e-7844ef645107%2Fa4486e7b-87d4-4354-a844-b53722d937ac%2Fe2865569-21ec-4213-8a1d-4f0ed5a1672e%2FCS_Logo_2022_In-Line_All-Red_RGB_(1).png%22%2C%22alt%22%3A%22CrowdStrike%20logo%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fct.capterra.com%2Fcapterra_tracker.gif%22%2C%22alt%22%3A%22%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fcontent%2Fdam%2Fcrowdstrike%2Fmarketing%2Fen-us%2Fimages%2Fadversaries%2Fadversary-group.png%22%2C%22alt%22%3A%22%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fcontent%2Fdam%2Fcrowdstrike%2Fmarketing%2Fen-us%2Fimages%2Fadversaries%2Fhi-res-adv-bg.png%22%2C%22alt%22%3A%22%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fcontent%2Fdam%2Fcrowdstrike%2Fmarketing%2Fen-us%2Fimages%2Fadversaries%2Fopen-report.png%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%5D%2C%22tier%22%3A%22PAID_QUOTA_TIER%22%2C%22pageUrl%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F%22%7D", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=c751ff21-98a4-40d0-886a-1297c969c40a&amp;amp;batch_time=1746729082514", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082089,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;9fdcab79-fdc2-4bb6-bc73-3be150907370&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/2a767faa4850d16352fcab54793a9929-20250227.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:98600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082089,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;67aacb3c-dd44-4d7f-ac20-42d0af89fcea&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/8286945929327c4f052d4cbfa8b74fed-20250403.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:102600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082089,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;ac02db19-cde7-4535-baf7-1dcf651cc1f8&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/a967dab279462cd1cebe42200b4f2f33-20250228.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:109300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082089,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;53f932b8-8e09-4979-b802-0f47ac8f0746&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/6ef1b26d9f40cf23fafeb151b569249b-20250501.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:110000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082089,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;c973556a-1811-4693-b784-c774954eaa0c&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/ab4b7b6f9a506e731d820031d0d59176-20250404.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:111200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082089,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;2f0e0cbb-1473-4b69-9ba1-85879c3b0b7e&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://d20vwa69zln1wj.cloudfront.net/7737a29b854de71521b1cd72c4118cfc/main/b47efe1afc5f0e84ad9f182e2e47eab7-20250228.jpg&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:118100000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082185,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;1f787744-647b-4051-8766-8b94a4f6607b&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://c.contentsquare.net/dvar?v=15.85.1&amp;pid=29632&amp;pn=1&amp;sn=1&amp;uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;dv=H4sIAAAAAAAAA6tWcvaIdwmKD8hJrEzKz89WcMssSk1RslJyTswtSMxMz1PwdLFSMLI0MrGwNFGqBQDRyonVLwAAAA%3D%3D&amp;ct=2&amp;r=868288&#34;,&#34;protocol&#34;:&#34;h2&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:83500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:0,&#34;encoded_body_size&#34;:0,&#34;decoded_body_size&#34;:0,&#34;transfer_size&#34;:300,&#34;download&#34;:{&#34;duration&#34;:0,&#34;start&#34;:83500000},&#34;first_byte&#34;:{&#34;duration&#34;:83300000,&#34;start&#34;:200000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082153,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;,&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;c84bc2fb-da94-442d-8d62-49a04fbc177e&#34;,&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;325f1466-d823-4b6f-8c51-4c6e4b232934&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:145000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://ws.zoominfo.com/pixel/collect&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;action&#34;:{&#34;target&#34;:{&#34;width&#34;:15,&#34;height&#34;:15,&#34;selector&#34;:&#34;#onetrust-close-btn-container&gt;BUTTON.onetrust-close-btn-handler&#34;},&#34;position&#34;:{&#34;x&#34;:8,&#34;y&#34;:7}}},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080587,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;,&#34;in_foreground&#34;:true},&#34;action&#34;:{&#34;id&#34;:&#34;0ad3f600-e8fc-42a9-b005-4e702367112a&#34;,&#34;target&#34;:{&#34;name&#34;:&#34;Close&#34;},&#34;type&#34;:&#34;click&#34;,&#34;loading_time&#34;:1711000000,&#34;frustration&#34;:{&#34;type&#34;:[&#34;error_click&#34;]},&#34;error&#34;:{&#34;count&#34;:1},&#34;long_task&#34;:{&#34;count&#34;:0},&#34;resource&#34;:{&#34;count&#34;:71}},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;type&#34;:&#34;action&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;action&#34;:{&#34;target&#34;:{&#34;width&#34;:15,&#34;height&#34;:15,&#34;selector&#34;:&#34;#onetrust-close-btn-container&gt;BUTTON.onetrust-close-btn-handler&#34;},&#34;position&#34;:{&#34;x&#34;:8,&#34;y&#34;:7}}},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729080669,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;,&#34;in_foreground&#34;:true},&#34;action&#34;:{&#34;id&#34;:&#34;282a7d4a-d70e-4ebb-84a0-31634014b1ea&#34;,&#34;target&#34;:{&#34;name&#34;:&#34;Close&#34;},&#34;type&#34;:&#34;click&#34;,&#34;loading_time&#34;:1629000000,&#34;frustration&#34;:{&#34;type&#34;:[&#34;error_click&#34;]},&#34;error&#34;:{&#34;count&#34;:1},&#34;long_task&#34;:{&#34;count&#34;:0},&#34;resource&#34;:{&#34;count&#34;:62}},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;type&#34;:&#34;action&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082430,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;589ef831-b643-4152-9beb-d36b439d021b&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/ipv?_biz_r=&amp;_biz_h=-1719904874&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079103&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=0&amp;a=crowdstrike.com&amp;rnd=480761&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:37200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082430,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;5813a49d-7105-4b0a-9d99-3eb8b13cbd40&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizibly.com/u?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079104&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;a=crowdstrike.com&amp;rnd=110587&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:38200000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082430,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;524c1329-998d-40e4-9786-0d5acca212c6&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/u?mapType=mkto&amp;mapValue=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729080105&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=1&amp;a=crowdstrike.com&amp;rnd=598845&amp;cdn_o=a&amp;_biz_z=1746729080105&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:38500000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082430,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;69207baf-5ac3-433e-973a-dac23d7bafe5&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/u?mapType=ecid&amp;mapValue=06D71E9261F941560A495CD6%40AdobeOrg_28180928486187528650392535887012291588&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729080106&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=2&amp;a=crowdstrike.com&amp;rnd=228061&amp;cdn_o=a&amp;_biz_z=1746729080208&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:40000000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082471,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6c13c131-9664-4f37-af37-a7905aebcf68&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/ipv?_biz_r=&amp;_biz_h=-1719904874&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079103&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=0&amp;a=crowdstrike.com&amp;rnd=480761&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:13300000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/snowflake/analyticsdata?timestamp=1746729082534", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;event-name&#34;:&#34;close cart&#34;,&#34;tag&#34;:&#34;Button&#34;,&#34;tag-location&#34;:&#34;Cart Icon&#34;,&#34;landingpage-url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;window-inner-width&#34;:1280,&#34;window-inner-height&#34;:720,&#34;body-client-width&#34;:1280,&#34;body-client-height&#34;:720,&#34;screen-width&#34;:1280,&#34;screen-height&#34;:720,&#34;screen-width-height&#34;:&#34;1280x720&#34;,&#34;device-pixel-ratio&#34;:1,&#34;browser-name&#34;:&#34;Chrome&#34;,&#34;browser-version&#34;:&#34;136.0.7103.25&#34;,&#34;os-name&#34;:&#34;Windows&#34;,&#34;os-version&#34;:&#34;NT 10.0&#34;,&#34;engine-name&#34;:&#34;Blink&#34;,&#34;engine-version&#34;:&#34;&#34;,&#34;device-type&#34;:&#34;desktop&#34;,&#34;session_id&#34;:&#34;8ffaa7ae-9b97-463d-b437-ce63a0483042&#34;,&#34;productName&#34;:&#34;&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://api.userway.org/api/br-links/v0/links/2376540", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://obs.fishrobotflower.com/mon", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`e=37dfbd8ee84e00126ee8c037e347829d9225c24f567d43d6da1908be6245cad7bd70a976710ce60ed89373bfe70e9c20c1e53e8d5a138c6c2717071a10acf9f29f671a82d589562b6d1ef878215780338a33c103300376c5065359600d5ccfbf394f77be26bb25cb43e2913df05365ac0b297d1bdc57e446f492d7da3dbb280ef67ccaa8073f8d0e3717224793845731f360b3f493a0180dec1edae97dfa2bc8169b1adc597cff3200e714561c4b92177af998ffe4198b6dec06c213f85e162ae7d133722b325f817c99ec59b058609fc6e359143e3dd385293e88864c06513c157a77bb9e70392652b48d1c2ad7f4ec3ee3b8192d4079b4a7a6928677a0dadf5ae8489e5d2e019cbecbf7af2b95dfe57594351ccdeb8b795904fd7367a095166dbdd0aa73d902fd12b0801d91499d9978953b67919d27dc6796d3bbe193fdbd4c38fc28b0bdfd354371fe8f719aa61af7010642dd4245c2979684c0ec8825ce3bc68ebe817fa27d763bc97f27db838003679ea875ac41bfc003d667a87e5346588c6eee5f8b89ee7ef8df78f02da7c19cdb27d60a0a2dea5d540b28073ed7ab67b355c4cb9764fdde25b3d02dd0c2b1be7c85b7f409746c89a91004d7b41d5107f7fae939fa2890f802fdfc3ce52dd36dcc068a50aefc2ecd0823e2ac90633f3d6402340a5b13ace68e65d205be5752517faac71386cd464887eccdb0ae6ddcece3f880ffadf84b4835fe4586a85e6b844a601fd16a59746ef07aa6771ec1f9dd71c83cc52fe0a82cc8a40e99721322210760129cc0e4db207a559a184689d2d9bbb0c24f049eb0be7f8ad2757adbce0f3a68fd8dc45020b23b98599e1cff22981ad8d40efbb337fe868bd49e5b7feee3337f9f71a68aacb6119d6d280b02956a5f4db6a25b93f112ca9887d552cd839b2e083745d321b951ef0ce6936384335d20463e7b78e3538e863a4bcf1cc7c2f03da8cb0be60d76e678072c8df4c3491fb725e78402db1930074eba3c2e8bafd061d52c49e4ff15ebbd9d88e11aa90924605fe06b3716d9c9584eb02ee6e0c02a8bd5eb0c547ca5633e507e867459e6c8e66c40cd260bbbd6c353d5c6f88f0f30b05bd61b781488173175d4cc1615da92b2aabde142a563f930b77096d65d9afff22a2bfdd215808cb4b1c4f6cadf741fddf42f54872dd257c81cf3ee28a8cad1c731efc4c7f2fb5a69cba2b5af51cb5594404921413d88823335dc60f64f8ac965ef0522e2714c9766e69c7575b97521525796f3d12f87c74c603d1e6ac2d583ffce74db&amp;cri=jfECv66HI9&amp;sf=0&amp;dc=&amp;cp=3&amp;gtm=WyJPbmVUcnVzdExvYWRlZCIsIk9wdGFub25Mb2FkZWQiLCJPbmVUcnVzdEdyb3Vwc1VwZGF0ZWQiLCJjb252ZXJzaW9uIiwicGFnZV92aWV3Il0%3D&amp;gac=462250241.1746729079&amp;mm=4%7C1115%2C518%2C0%2C0%2C1115%2C518%2C518%2C1%2C1115%2C518%7C640%2C142%2C-277%2C108%2C640%2C235%2C2465%2C1%2C640%2C142&amp;md=4%7C1115%2C518%2C1115%2C518%2C520%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C640%2C142%2C640%2C235%2C2466%2C1%2C%2C%2C640%2C142&amp;mu=4%7C1115%2C518%2C1115%2C518%2C521%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C640%2C142%2C640%2C235%2C2467%2C1%2C%2C%2C640%2C142&amp;cl=4%7C1115%2C518%2C1115%2C518%2C528%2C1%2CYnV0dG9uLDE1LDE1LDExMDcsNTExLDAsMA%3D%3D%2C%2C1115%2C518%7C640%2C142%2C640%2C235%2C2468%2C1%2C%2C%2C640%2C142&amp;tb=1&amp;gi=9%2C2448%7C12%2C2449%7C10%2C2450%7C14%2C2452%7C11%2C2452%7C13%2C2452%7C21%2C2461%7C18%2C2862&amp;ws=1280x720&amp;wos=1280x720&amp;ver=13&amp;fi=517&amp;ti=3015&amp;mo=0&amp;pn=4946&amp;spn=1931&amp;fp=516&amp;snt=1&amp;sc=10%7C0%2C2%2C931%2C1%7C0%2C93%2C1080%2C1`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://c.contentsquare.net/v2/events?uu=e032267c-9562-a268-d804-cdcd5a5b9f8e&amp;amp;sn=1&amp;amp;hd=1746729079&amp;amp;v=15.85.1&amp;amp;pid=29632&amp;amp;pn=1&amp;amp;sr=10&amp;amp;mdh=7994&amp;amp;str=96&amp;amp;di=485&amp;amp;dc=2326&amp;amp;fl=2347&amp;amp;ct=0", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`[{&#34;type&#34;:11,&#34;ts&#34;:1922,&#34;tgt&#34;:&#34;div#searchfield&gt;div:eq(0)&gt;form:eq(0)&gt;div:eq(0)&gt;input:eq(0)&#34;},{&#34;type&#34;:7,&#34;ts&#34;:1928,&#34;x&#34;:917,&#34;y&#34;:127,&#34;tgt&#34;:&#34;div#search-header&gt;button:eq(0)&gt;img:eq(0)&#34;},{&#34;type&#34;:2,&#34;ts&#34;:2815,&#34;x&#34;:640,&#34;y&#34;:235,&#34;xRel&#34;:32768,&#34;yRel&#34;:32070,&#34;tgtHM&#34;:&#34;div#results&gt;div:eq(0)&gt;div:eq(0)&#34;},{&#34;type&#34;:3,&#34;ts&#34;:2816,&#34;x&#34;:640,&#34;y&#34;:235,&#34;tgt&#34;:&#34;div#results&gt;div:eq(0)&gt;div:eq(0)&#34;},{&#34;type&#34;:10,&#34;ts&#34;:2816,&#34;tgt&#34;:&#34;input#addsearchfield&#34;,&#34;isBlank&#34;:false},{&#34;type&#34;:12,&#34;ts&#34;:2817,&#34;tgt&#34;:&#34;input#addsearchfield&#34;},{&#34;type&#34;:4,&#34;ts&#34;:2818,&#34;x&#34;:640,&#34;y&#34;:235,&#34;tgt&#34;:&#34;div#results&gt;div:eq(0)&gt;div:eq(0)&#34;},{&#34;type&#34;:5,&#34;ts&#34;:2818,&#34;x&#34;:640,&#34;y&#34;:235,&#34;tgt&#34;:&#34;div#results&gt;div:eq(0)&gt;div:eq(0)&#34;,&#34;xRel&#34;:32768,&#34;yRel&#34;:32070}]`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://px.ads.linkedin.com/wa/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;pids&#34;:[64444],&#34;scriptVersion&#34;:213,&#34;time&#34;:1746729083157,&#34;domain&#34;:&#34;crowdstrike.com&#34;,&#34;url&#34;:&#34;https://crowdstrike.com/en-us/?search=blash&#34;,&#34;pageTitle&#34;:&#34;CrowdStrike: We Stop Breaches with AI-native Cybersecurity&#34;,&#34;websiteSignalRequestId&#34;:&#34;03888bb6-d414-643f-b54e-849053480c93&#34;,&#34;isTranslated&#34;:false,&#34;liFatId&#34;:&#34;&#34;,&#34;liGiant&#34;:&#34;&#34;,&#34;misc&#34;:{&#34;psbState&#34;:-4},&#34;isLinkedInApp&#34;:false,&#34;hem&#34;:null,&#34;signalType&#34;:&#34;PAGE_VISIT&#34;}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.facebook.com/tr/?id=950083805267950&amp;amp;ev=PageView&amp;amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F%3Fsearch%3Dblash&amp;amp;rl=&amp;amp;if=false&amp;amp;ts=1746729083155&amp;amp;sw=1280&amp;amp;sh=720&amp;amp;v=2.9.201&amp;amp;r=stable&amp;amp;a=adobe_launch&amp;amp;ec=1&amp;amp;o=28&amp;amp;ler=empty&amp;amp;it=1746729079168&amp;amp;coo=false&amp;amp;exp=k2&amp;amp;rqm=GET", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://www.facebook.com/tr/?id=992980065451679&amp;amp;ev=PageView&amp;amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F%3Fsearch%3Dblash&amp;amp;rl=&amp;amp;if=false&amp;amp;ts=1746729083157&amp;amp;sw=1280&amp;amp;sh=720&amp;amp;v=2.9.201&amp;amp;r=stable&amp;amp;ec=3&amp;amp;o=4126&amp;amp;fbp=fb.1.1746729079520.75795469452619173&amp;amp;ler=empty&amp;amp;it=1746729079168&amp;amp;coo=false&amp;amp;exp=k2&amp;amp;rqm=GET", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://browser-intake-datadoghq.com/api/v2/rum?ddsource=browser&amp;amp;ddtags=sdk_version%3A5.35.1%2Capi%3Afetch%2Cenv%3Aproduction-aem%2Cservice%3Amarketing-web&amp;amp;dd-api-key=pub706deb6cb525521c08e42ee8885babbe&amp;amp;dd-evp-origin-version=5.35.1&amp;amp;dd-evp-origin=browser&amp;amp;dd-request-id=2edef67a-3a74-4dd8-9815-13bb59694efa&amp;amp;batch_time=1746729083180", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082471,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;a92d47bf-6a80-467b-b324-4230ff5a4257&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizibly.com/u?_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729079104&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;a=crowdstrike.com&amp;rnd=110587&amp;cdn_o=a&amp;_biz_z=1746729079104&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:16600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082471,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;918b3132-c97b-465b-b2ad-0d66460e3e08&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/u?mapType=mkto&amp;mapValue=id%3A281-OBQ-266%26token%3A_mch-crowdstrike.com-d292972b603d4598c10cdd175d1f8c61&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729080105&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=1&amp;a=crowdstrike.com&amp;rnd=598845&amp;cdn_o=a&amp;_biz_z=1746729080105&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:17700000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082471,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;7fe26532-7036-4ece-86d8-41ad699dcb1e&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;70be626f-db4e-4270-b25a-6eadfebfd656&#34;,&#34;type&#34;:&#34;image&#34;,&#34;url&#34;:&#34;https://cdn.bizible.com/u?mapType=ecid&amp;mapValue=06D71E9261F941560A495CD6%40AdobeOrg_28180928486187528650392535887012291588&amp;_biz_u=35a249abd9e94512be1326ac06e4ddd5&amp;_biz_l=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;_biz_t=1746729080106&amp;_biz_i=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;_biz_n=2&amp;a=crowdstrike.com&amp;rnd=228061&amp;cdn_o=a&amp;_biz_z=1746729080208&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;duration&#34;:18600000,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;action&#34;:{&#34;target&#34;:{&#34;width&#34;:20,&#34;height&#34;:20,&#34;selector&#34;:&#34;#search-header&gt;BUTTON.search_btn&gt;IMG&#34;},&#34;position&#34;:{&#34;x&#34;:10,&#34;y&#34;:10}}},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729081617,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;,&#34;in_foreground&#34;:true},&#34;action&#34;:{&#34;id&#34;:&#34;df429e33-fde0-41d5-aa92-322f8a164772&#34;,&#34;target&#34;:{&#34;name&#34;:&#34;Search Icon&#34;},&#34;type&#34;:&#34;click&#34;,&#34;loading_time&#34;:681000000,&#34;frustration&#34;:{&#34;type&#34;:[]},&#34;error&#34;:{&#34;count&#34;:0},&#34;long_task&#34;:{&#34;count&#34;:0},&#34;resource&#34;:{&#34;count&#34;:18}},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;type&#34;:&#34;action&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082535,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;bdaac348-3f9d-45bf-91e3-0c1af40a5b9b&#34;,&#34;99b5b350-3db2-43ca-9d8e-3fbee2d04f2b&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;b90f7a34-ff17-413f-a8b7-57bcdd7f77d6&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:152300000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;protocol&#34;:&#34;h2&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/bin/crowdstrike/nativeshopping/v1/snowflake/analyticsdata?timestamp=1746729082534&#34;,&#34;delivery_type&#34;:&#34;other&#34;,&#34;render_blocking_status&#34;:&#34;non-blocking&#34;,&#34;size&#34;:153,&#34;encoded_body_size&#34;:127,&#34;decoded_body_size&#34;:153,&#34;transfer_size&#34;:427,&#34;download&#34;:{&#34;duration&#34;:400000,&#34;start&#34;:151900000},&#34;first_byte&#34;:{&#34;duration&#34;:151600000,&#34;start&#34;:300000}},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082473,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;a1750461-0d4a-4272-8e0b-b7cff625edc9&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:374000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://cdn77.api.userway.org/api/img-dscr/v2/dyvvHf6oG0/2376540/V4sorvbtJTNsh59B/alts.json?dto=%7B%22sorted%22%3A%5B%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fblack-primary-crowdstrike-logo-1-addedPadding-3%3Fts%3D1746646455211%26dpr%3Doff%22%2C%22alt%22%3A%22CrowdStrike%20Logo%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FCloud-Security-2%22%2C%22alt%22%3A%22Cloud%20Security%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FData-Protection-3%22%2C%22alt%22%3A%22Data%20Protection%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fempty-cart-image%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FEndpoint-Security-2%22%2C%22alt%22%3A%22Endpoint%20Security%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fexperience-breach-Icon%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FExposure-Management-2%22%2C%22alt%22%3A%22Exposure%20Management%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FGenerative-AI%22%2C%22alt%22%3A%22Generative%20AI%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FIdentity-Protection-1%22%2C%22alt%22%3A%22Identity%20Protection%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FIT-Automation%22%2C%22alt%22%3A%22IT%20Automation%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Flogin-icon%22%2C%22alt%22%3A%22Login%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FNext-Gen-SIEM%22%2C%22alt%22%3A%22Next-Gen%20SIEM%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fphone-icon%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FPlatform-4-16-25-darkmode-HP%3Fts%3D1746621981798%26dpr%3Doff%22%2C%22alt%22%3A%22platform%20graphic%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fprivacyoptions-icon-footer%22%2C%22alt%22%3A%22Your%20Privacy%20Choices%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FSaaS-Security%22%2C%22alt%22%3A%22SaaS%20Security%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fshopping-cart-empty%22%2C%22alt%22%3A%22cart%20icon%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FThreat-Intelligence-Hunting%22%2C%22alt%22%3A%22Threat%20Intelligence%20%26%20Hunting%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2FWorkflow-Automation%22%2C%22alt%22%3A%22Workflow%20Automation%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fcontent%2Fcrowdstrikeinc%2Fxiot-security%22%2C%22alt%22%3A%22XIoT%20Security%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2F25-SRV-007_Forrester%2520Wave%2520MDR%25203153x3860%3Fts%3D1746621981959%26dpr%3Doff%22%2C%22alt%22%3A%22Forrester%20MDR%202025%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fcharlotte-hp-hero-1%3F%26hei%3D769%26wid%3D1600%26qlt%3D90%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fcloud-lock-generic-img%3Fts%3D1746621982750%26dpr%3Doff%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Ffrost-radar-2024-2%3Fts%3D1746621982168%26dpr%3Doff%22%2C%22alt%22%3A%22Frost%20CNAPP%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fgartner-mq-leader-report%3Fts%3D1746621982064%26dpr%3Doff%22%2C%22alt%22%3A%22Gigaom%20Ransomware%20Protection%20Chart%202024%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fidentity-security-services-new%3Fts%3D1746621982825%26dpr%3Doff%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Fsoc-survival-guide-2.bak%3Fts%3D1746621982899%26dpr%3Doff%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fassets.crowdstrike.com%2Fis%2Fimage%2Fcrowdstrikeinc%2Ftry-free-footer%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fcdn.cookielaw.org%2Flogos%2Fc109dae9-46f3-4e91-a59e-7844ef645107%2Fa4486e7b-87d4-4354-a844-b53722d937ac%2Fe2865569-21ec-4213-8a1d-4f0ed5a1672e%2FCS_Logo_2022_In-Line_All-Red_RGB_(1).png%22%2C%22alt%22%3A%22CrowdStrike%20logo%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fct.capterra.com%2Fcapterra_tracker.gif%22%2C%22alt%22%3A%22%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fcontent%2Fdam%2Fcrowdstrike%2Fmarketing%2Fen-us%2Fimages%2Fadversaries%2Fadversary-group.png%22%2C%22alt%22%3A%22%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fcontent%2Fdam%2Fcrowdstrike%2Fmarketing%2Fen-us%2Fimages%2Fadversaries%2Fhi-res-adv-bg.png%22%2C%22alt%22%3A%22%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fcontent%2Fdam%2Fcrowdstrike%2Fmarketing%2Fen-us%2Fimages%2Fadversaries%2Fopen-report.png%22%2C%22alt%22%3A%22%22%2C%22dir%22%3A%22RO%22%7D%5D%2C%22tier%22%3A%22PAID_QUOTA_TIER%22%2C%22pageUrl%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F%22%7D&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729082855,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;bdaac348-3f9d-45bf-91e3-0c1af40a5b9b&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;6d5f15d7-72a6-433e-a676-8530399f3019&#34;,&#34;type&#34;:&#34;fetch&#34;,&#34;duration&#34;:52000000,&#34;method&#34;:&#34;GET&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://api.userway.org/api/br-links/v0/links/2376540&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:0,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0},&#34;discarded&#34;:false},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729083082,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;},&#34;action&#34;:{&#34;id&#34;:[&#34;bdaac348-3f9d-45bf-91e3-0c1af40a5b9b&#34;]},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;resource&#34;:{&#34;id&#34;:&#34;95f62e5f-518f-42b3-be49-f408cca09236&#34;,&#34;type&#34;:&#34;xhr&#34;,&#34;duration&#34;:83000000,&#34;method&#34;:&#34;POST&#34;,&#34;status_code&#34;:200,&#34;url&#34;:&#34;https://obs.fishrobotflower.com/mon&#34;},&#34;type&#34;:&#34;resource&#34;}
{&#34;_dd&#34;:{&#34;format_version&#34;:2,&#34;drift&#34;:-1,&#34;configuration&#34;:{&#34;session_sample_rate&#34;:100,&#34;session_replay_sample_rate&#34;:0,&#34;start_session_replay_recording_manually&#34;:true},&#34;document_version&#34;:2,&#34;page_states&#34;:[{&#34;state&#34;:&#34;active&#34;,&#34;start&#34;:1615600000}]},&#34;application&#34;:{&#34;id&#34;:&#34;2f31f954-44f4-405b-9bdf-b2a0775f63b4&#34;},&#34;date&#34;:1746729078135,&#34;service&#34;:&#34;marketing-web&#34;,&#34;source&#34;:&#34;browser&#34;,&#34;session&#34;:{&#34;id&#34;:&#34;106ffeda-51bb-4c48-9d16-4ee3a0f882b8&#34;,&#34;type&#34;:&#34;user&#34;,&#34;sampled_for_replay&#34;:false},&#34;view&#34;:{&#34;id&#34;:&#34;fabbc0ae-b947-4067-81a1-b92ef25d8ce8&#34;,&#34;url&#34;:&#34;https://www.crowdstrike.com/en-us/&#34;,&#34;referrer&#34;:&#34;&#34;,&#34;action&#34;:{&#34;count&#34;:3},&#34;frustration&#34;:{&#34;count&#34;:2},&#34;cumulative_layout_shift&#34;:0.7721,&#34;cumulative_layout_shift_time&#34;:583600000,&#34;cumulative_layout_shift_target_selector&#34;:&#34;BODY&gt;DIV.root&gt;DIV.cmp-container&gt;DIV.aem-Grid&gt;DIV.experiencefragment&gt;DIV.cmp-experiencefragment&gt;DIV.cmp-container&gt;DIV.aem-Grid&gt;DIV.container&gt;DIV.cmp-container&gt;DIV.container&gt;DIV.cmp-container&gt;DIV.text&#34;,&#34;first_byte&#34;:152200000,&#34;dom_complete&#34;:2384400000,&#34;dom_content_loaded&#34;:569500000,&#34;dom_interactive&#34;:543500000,&#34;error&#34;:{&#34;count&#34;:1},&#34;first_contentful_paint&#34;:516000000,&#34;first_input_delay&#34;:1300000,&#34;first_input_time&#34;:2448100000,&#34;first_input_target_selector&#34;:&#34;#onetrust-close-btn-container&gt;BUTTON.onetrust-close-btn-handler&#34;,&#34;interaction_to_next_paint&#34;:48000000,&#34;interaction_to_next_paint_time&#34;:2448100000,&#34;interaction_to_next_paint_target_selector&#34;:&#34;#onetrust-close-btn-container&gt;BUTTON.onetrust-close-btn-handler&#34;,&#34;is_active&#34;:true,&#34;largest_contentful_paint&#34;:1116000000,&#34;largest_contentful_paint_target_selector&#34;:&#34;#homepage-hero-carousel&gt;DIV.cmp-hero-carousel__content&#34;,&#34;load_event&#34;:2405800000,&#34;loading_time&#34;:4163000000,&#34;loading_type&#34;:&#34;initial_load&#34;,&#34;long_task&#34;:{&#34;count&#34;:2},&#34;resource&#34;:{&#34;count&#34;:248},&#34;time_spent&#34;:4619000000},&#34;display&#34;:{&#34;viewport&#34;:{&#34;width&#34;:1280,&#34;height&#34;:720},&#34;scroll&#34;:{&#34;max_depth&#34;:813,&#34;max_depth_scroll_top&#34;:93,&#34;max_scroll_height&#34;:7994,&#34;max_scroll_height_time&#34;:2630000000}},&#34;connectivity&#34;:{&#34;status&#34;:&#34;connected&#34;,&#34;effective_type&#34;:&#34;4g&#34;},&#34;type&#34;:&#34;view&#34;,&#34;privacy&#34;:{&#34;replay_level&#34;:&#34;mask-user-input&#34;}}`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		t.Errorf("Expected status %d, got %d", 202, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://metrics.api.drift.com/monitoring/metrics/event3/bulk", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`metric_payload=%7B%22id%22%3A%22fa8f1e52-a184-472e-9268-f22272c10906%22%2C%22events%22%3A%5B%7B%22eventName%22%3A%22%5BV2%5D%20-%20enduser_bootstrap%22%2C%22orgId%22%3A113673%2C%22embedId%22%3A%229d4udx6ceimp%22%2C%22sessionId%22%3A%22bb076d85-176d-40bf-8edd-e4d09c550682%22%2C%22instanceId%22%3A%225e2d5498-aa99-4851-a408-bdc0dce124ee%22%2C%22attributes%22%3A%7B%22duration%22%3A517.5%2C%22isMobile%22%3Afalse%7D%7D%2C%7B%22eventName%22%3A%22%5BV2%5D%20-%20loaded%22%2C%22orgId%22%3A113673%2C%22embedId%22%3A%229d4udx6ceimp%22%2C%22sessionId%22%3A%22bb076d85-176d-40bf-8edd-e4d09c550682%22%2C%22instanceId%22%3A%225e2d5498-aa99-4851-a408-bdc0dce124ee%22%2C%22attributes%22%3A%7B%22duration%22%3A1429.5999999940395%2C%22isMobile%22%3Afalse%2C%22loadSizeInBytes%22%3A70112%7D%7D%5D%2C%22timestamp%22%3A1746729083528%2C%22context%22%3A%7B%22country%22%3A%22US%22%2C%22locale%22%3A%22en-US%22%2C%22timezone%22%3A%22America%2FLos_Angeles%22%2C%22url%22%3A%22crowdstrike.com%2Fen-us%22%7D%7D`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test POST request
	req, err = http.NewRequest("POST", "https://www.google-analytics.com/g/collect?v=2&amp;amp;tid=G-ZKTET1D58V&amp;amp;gtm=45je5570h1v894068940za200&amp;amp;_p=1746729079180&amp;amp;gcs=G111&amp;amp;gcd=13n3n3m3m5l1&amp;amp;npa=1&amp;amp;dma_cps=-&amp;amp;dma=0&amp;amp;tag_exp=101509157~103101750~103101752~103116025~103200001~103233427~103251618~103251620~103284320~103284322~103301114~103301116&amp;amp;gdid=dYWJhMj&amp;amp;cid=462250241.1746729079&amp;amp;ul=en-us&amp;amp;are=1&amp;amp;frm=0&amp;amp;pscdl=noapi&amp;amp;_eu=AAAAAAQ&amp;amp;_geo=1&amp;amp;_rdi=1&amp;amp;sid=1746729079&amp;amp;sct=1&amp;amp;seg=1&amp;amp;dl=https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F&amp;amp;dt=CrowdStrike%3A%20We%20Stop%20Breaches%20with%20AI-native%20Cybersecurity&amp;amp;_s=2&amp;amp;tfd=6035", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	
	req.Body = io.NopCloser(bytes.NewBufferString(`en=page_view&amp;_ee=1&amp;ep.additional_comment=Company%20name%20or%20domain%20match%20was%20found&amp;ep.address=7509%20E%201ST%20ST&amp;ep.annual_revenue=5772000&amp;ep.city=Scottsdale&amp;ep.company_match=Match&amp;ep.country=United%20States&amp;ep.country_iso_code=US&amp;ep.domain=burtfeldman.com&amp;ep.employee_count=1&amp;ep.employee_range=0%20-%209&amp;ep.geoIP_city=Sunnyvale&amp;ep.geoIP_country=United%20States&amp;ep.geoIP_state=California&amp;ep.industry=Business%20Services&amp;ep.is_6qa=false&amp;ep.is_blacklisted=false&amp;ep.naics=541110&amp;ep.naics_description=Offices%20of%20Lawyers&amp;ep.name=Burt%20Feldman&amp;ep.region=Northern%20America&amp;ep.revenue_range=%245M%20-%20%2410M&amp;ep.sic=8111&amp;ep.sic_description=Legal%20Services&amp;ep.state=Arizona&amp;ep.state_code=AZ&amp;ep.zip=85251&amp;ep.confidence=Low&amp;epn.segments_ids=386105&amp;ep.industry_v2_industry=Business%20Services&amp;ep.industry_v2_subindustry=Legal&amp;_et=2015
en=page_view&amp;_ee=1&amp;ep.additional_comment=Company%20name%20or%20domain%20match%20was%20found&amp;ep.address=7509%20E%201ST%20ST&amp;ep.annual_revenue=5772000&amp;ep.city=Scottsdale&amp;ep.company_match=Match&amp;ep.country=United%20States&amp;ep.country_iso_code=US&amp;ep.domain=burtfeldman.com&amp;ep.employee_count=1&amp;ep.employee_range=0%20-%209&amp;ep.geoIP_city=Sunnyvale&amp;ep.geoIP_country=United%20States&amp;ep.geoIP_state=California&amp;ep.industry=Business%20Services&amp;ep.is_6qa=false&amp;ep.is_blacklisted=false&amp;ep.naics=541110&amp;ep.naics_description=Offices%20of%20Lawyers&amp;ep.name=Burt%20Feldman&amp;ep.region=Northern%20America&amp;ep.revenue_range=%245M%20-%20%2410M&amp;ep.sic=8111&amp;ep.sic_description=Legal%20Services&amp;ep.state=Arizona&amp;ep.state_code=AZ&amp;ep.zip=85251&amp;ep.confidence=Low&amp;epn.segments_ids=386105&amp;ep.industry_v2_industry=Business%20Services&amp;ep.industry_v2_subindustry=Legal&amp;_et=1174
en=page_view&amp;_ee=1&amp;ep.additional_comment=Company%20name%20or%20domain%20match%20was%20found&amp;ep.address=7509%20E%201ST%20ST&amp;ep.annual_revenue=5772000&amp;ep.city=Scottsdale&amp;ep.company_match=Match&amp;ep.country=United%20States&amp;ep.country_iso_code=US&amp;ep.domain=burtfeldman.com&amp;ep.employee_count=1&amp;ep.employee_range=0%20-%209&amp;ep.geoIP_city=Sunnyvale&amp;ep.geoIP_country=United%20States&amp;ep.geoIP_state=California&amp;ep.industry=Business%20Services&amp;ep.is_6qa=false&amp;ep.is_blacklisted=false&amp;ep.naics=541110&amp;ep.naics_description=Offices%20of%20Lawyers&amp;ep.name=Burt%20Feldman&amp;ep.region=Northern%20America&amp;ep.revenue_range=%245M%20-%20%2410M&amp;ep.sic=8111&amp;ep.sic_description=Legal%20Services&amp;ep.state=Arizona&amp;ep.state_code=AZ&amp;ep.zip=85251&amp;ep.confidence=Low&amp;epn.segments_ids=386105&amp;ep.industry_v2_industry=Business%20Services&amp;ep.industry_v2_subindustry=Legal&amp;_et=4`))
	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Expected status %d, got %d", 204, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://cdn77.api.userway.org/api/img-dscr/v2/dyvvHf6oG0/2376540/V4sorvbtJTNsh59B/alts.json?dto=%7B%22sorted%22%3A%5B%7B%22src%22%3A%22https%3A%2F%2Fd20vwa69zln1wj.cloudfront.net%2F7737a29b854de71521b1cd72c4118cfc%2Fmain%2F2a767faa4850d16352fcab54793a9929-20250227.jpg%22%2C%22alt%22%3A%22How%20Falcon%20OverWatch%20Hunts%20for%20Out-of-Band%20Application%20Security%20Testing%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fd20vwa69zln1wj.cloudfront.net%2F7737a29b854de71521b1cd72c4118cfc%2Fmain%2F2a767faa4850d16352fcab54793a9929-20250302.jpg%22%2C%22alt%22%3A%22CrowdStrike%20Defends%20Against%20Azure%20Cross-Tenant%20Synchronization%20Attacks%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fd20vwa69zln1wj.cloudfront.net%2F7737a29b854de71521b1cd72c4118cfc%2Fmain%2F6ef1b26d9f40cf23fafeb151b569249b-20250501.jpg%22%2C%22alt%22%3A%22Accelerate%20Investigations%20With%20AI%20%7C%20CrowdStrike%20Solutions%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fd20vwa69zln1wj.cloudfront.net%2F7737a29b854de71521b1cd72c4118cfc%2Fmain%2F8286945929327c4f052d4cbfa8b74fed-20250403.jpg%22%2C%22alt%22%3A%22What%20is%20Zero%20Trust%3F%20-%20Guide%20to%20Zero%20Trust%20Security%20%7C%20CrowdStrike%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fd20vwa69zln1wj.cloudfront.net%2F7737a29b854de71521b1cd72c4118cfc%2Fmain%2Fa967dab279462cd1cebe42200b4f2f33-20250228.jpg%22%2C%22alt%22%3A%22How%20to%20Harden%20Your%20Cloud%20Against%20SMTP%20Abuse%20%7C%20CrowdStrike%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fd20vwa69zln1wj.cloudfront.net%2F7737a29b854de71521b1cd72c4118cfc%2Fmain%2Fab4b7b6f9a506e731d820031d0d59176-20250404.jpg%22%2C%22alt%22%3A%22Cushman%20%26%20Wakefield%3A%20Minimizing%20the%20Blast%20Radius%20of%20Security%20Events%22%2C%22dir%22%3A%22RO%22%7D%2C%7B%22src%22%3A%22https%3A%2F%2Fd20vwa69zln1wj.cloudfront.net%2F7737a29b854de71521b1cd72c4118cfc%2Fmain%2Fb47efe1afc5f0e84ad9f182e2e47eab7-20250228.jpg%22%2C%22alt%22%3A%22Microsoft%20Active%20Directory%20Supply%20Chain%20Compromise%20Reflects%20Shifting%20Adversary%20Tactics%20to%20Exploit%20Identity%22%2C%22dir%22%3A%22RO%22%7D%5D%2C%22tier%22%3A%22PAID_QUOTA_TIER%22%2C%22pageUrl%22%3A%22https%3A%2F%2Fwww.crowdstrike.com%2Fen-us%2F%3Fsearch%3Dblash%22%7D", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
	// Test GET request
	req, err = http.NewRequest("GET", "https://api.userway.org/api/br-links/v0/links/2376540", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	

	
	req.Header.Set("Accept-Language", "en-US")
	

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}
	
}
