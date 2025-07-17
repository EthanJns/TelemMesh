# lib/telem_dash_web/router.ex

defmodule TelemDashWeb.Router do
  use TelemDashWeb, :router

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_live_flash
    plug :put_root_layout, html: {TelemDashWeb.Layouts, :root}
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/", TelemDashWeb do
    pipe_through :browser

    get "/", PageController, :home
    live "/telemetry", TelemetryLive  # This is where we route to the LiveView
  end

  # Other routes here...

  scope "/api", TelemDashWeb do
    pipe_through :api
    post "/telemetry", TelemetryController, :create
  end
end
