defmodule TelemDash.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      TelemDashWeb.Telemetry,
      TelemDash.Repo,
      {DNSCluster, query: Application.get_env(:telem_dash, :dns_cluster_query) || :ignore},
      {Phoenix.PubSub, name: TelemDash.PubSub},
      {TelemDash.TelemetryLogs, []},
      # Start the Finch HTTP client for sending emails
      {Finch, name: TelemDash.Finch},
      # Start a worker by calling: TelemDash.Worker.start_link(arg)
      # {TelemDash.Worker, arg},
      # Start to serve requests, typically the last entry
      TelemDashWeb.Endpoint
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: TelemDash.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    TelemDashWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
