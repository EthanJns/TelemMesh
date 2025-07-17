defmodule TelemDashWeb.TelemetryController do
  use TelemDashWeb, :controller

  require Logger

  # --------------------------------------------------------------
  def create(conn, %{
    "node_id" => node_id,
    "data" => data,
    "timestamp_unix" => timestamp_unix
  }) do
    # Log the incoming telemetry data for debugging
    Logger.info("Received telemetry from #{node_id}: #{inspect(data)} at timestamp #{timestamp_unix}")

    # Broadcast the telemetry data to the LiveView
    Phoenix.PubSub.broadcast(
      TelemDash.PubSub,
      "telemetry_data",
      {:telemetry_data, %{node_id: node_id, data: data, timestamp_unix: timestamp_unix}}
    )

    send_resp(conn, 204, "")
  end

  # --------------------------------------------------------------
  def create(conn, _params) do
    send_resp(conn, 400, "Invalid telemetry data")
  end
end
