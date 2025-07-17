defmodule TelemDash.TelemetryLogs do
  use Agent

  # Starts the agent with an empty list to store logs
  def start_link(_opts) do
    Agent.start_link(fn -> [] end, name: __MODULE__)
  end

  # Adds a new log entry to the list
  def add_log(log) do
    Agent.update(__MODULE__, fn logs -> [log | logs] end)
  end

  # Retrieves the most recent logs (you can adjust the number here)
  def list_recent_logs do
    Agent.get(__MODULE__, fn logs -> logs end)
  end
end
