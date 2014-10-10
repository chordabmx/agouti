package mocks

import (
	"github.com/sclevine/agouti/page/internal/webdriver"
	"io"
)

type Driver struct {
	NavigateCall struct {
		URL string
		Err error
	}

	GetElementsCall struct {
		Selector       string
		ReturnElements []webdriver.Element
		Err            error
	}

	GetWindowCall struct {
		ReturnWindow webdriver.Window
		Err          error
	}

	ScreenshotCall struct {
		ReturnImage io.Reader
		Err         error
	}

	SetCookieCall struct {
		Cookie *webdriver.Cookie
		Err    error
	}

	DeleteCookieCall struct {
		Name string
		Err  error
	}

	DeleteAllCookiesCall struct {
		WasCalled bool
		Err       error
	}

	GetURLCall struct {
		ReturnURL string
		Err       error
	}
}

func (d *Driver) Navigate(url string) error {
	d.NavigateCall.URL = url
	return d.NavigateCall.Err
}

func (d *Driver) GetElements(selector string) ([]webdriver.Element, error) {
	d.GetElementsCall.Selector = selector
	return d.GetElementsCall.ReturnElements, d.GetElementsCall.Err
}

func (d *Driver) GetWindow() (webdriver.Window, error) {
	return d.GetWindowCall.ReturnWindow, d.GetWindowCall.Err
}

func (d *Driver) Screenshot() (io.Reader, error) {
	return d.ScreenshotCall.ReturnImage, d.ScreenshotCall.Err
}

func (d *Driver) SetCookie(cookie *webdriver.Cookie) error {
	d.SetCookieCall.Cookie = cookie
	return d.SetCookieCall.Err
}

func (d *Driver) DeleteCookie(name string) error {
	d.DeleteCookieCall.Name = name
	return d.DeleteCookieCall.Err
}

func (d *Driver) DeleteAllCookies() error {
	d.DeleteAllCookiesCall.WasCalled = true
	return d.DeleteAllCookiesCall.Err
}

func (d *Driver) GetURL() (string, error) {
	return d.GetURLCall.ReturnURL, d.GetURLCall.Err
}
