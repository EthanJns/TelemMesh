defmodule TelemDashWeb.TelemetryLive do
  use Phoenix.LiveView

  @topic "telemetry_data"

  def mount(_params, _session, socket) do
    if connected?(socket) do
      Phoenix.PubSub.subscribe(TelemDash.PubSub, @topic)
      IO.inspect(self(), label: "Subscribed to telemetry_data topic")
    end

    {:ok, assign(socket, telemetry_logs: [])}
  end

  def handle_info({:telemetry_data, data}, socket) do
    IO.inspect(self(), label: "LiveView PID")
    IO.inspect(data, label: "Received in LiveView")

    # Append data to telemetry_logs
    logs = [data | socket.assigns.telemetry_logs]
    {:noreply, assign(socket, telemetry_logs: logs)}
  end

  def render(assigns) do
  ~H"""
  <div class="min-h-screen bg-gray-900 text-white p-6">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-4xl font-bold text-blue-400 mb-2">Telemetry Dashboard</h1>
      <div class="flex items-center space-x-4 text-gray-300">
        <span class="bg-green-500 w-3 h-3 rounded-full animate-pulse"></span>
        <span>Live Data Stream</span>
        <span class="text-sm">Total Logs: <%= length(@telemetry_logs) %></span>
      </div>
    </div>

    <!-- Current Metrics Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
      <%= if length(@telemetry_logs) > 0 do %>
        <%= for log <- Enum.take(@telemetry_logs, 1) do %>
          <%= for metric <- log.data do %>
            <div class="bg-gray-800 rounded-lg p-6 border border-gray-700 hover:border-blue-500 transition-colors">
              <div class="flex items-center justify-between mb-4">
                <h3 class="text-lg font-semibold text-gray-200"><%= String.replace(metric["name"], "_", " ") |> String.capitalize() %></h3>
                <span class="text-xs text-gray-500">Node: <%= log.node_id %></span>
              </div>
              <div class="flex items-end space-x-2">
                <span class="text-3xl font-bold text-blue-400"><%= :erlang.float_to_binary(metric["value"], decimals: 2) %></span>
                <span class="text-sm text-gray-400 mb-1"><%= metric["unit"] %></span>
              </div>
              <div class="mt-4 text-xs text-gray-500">
                Last updated: <%= DateTime.from_unix!(log.timestamp_unix) |> DateTime.to_string() %>
              </div>
            </div>
          <% end %>
        <% end %>
      <% else %>
        <div class="col-span-full bg-gray-800 rounded-lg p-8 border border-gray-700 text-center">
          <div class="text-gray-400 mb-2">
            <svg class="w-12 h-12 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-300 mb-2">No Telemetry Data</h3>
          <p class="text-gray-500">Waiting for incoming telemetry data...</p>
        </div>
      <% end %>
    </div>

    <!-- Recent Activity Log -->
    <div class="bg-gray-800 rounded-lg border border-gray-700">
      <div class="p-6 border-b border-gray-700">
        <h2 class="text-xl font-semibold text-gray-200 flex items-center">
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          Recent Activity
        </h2>
      </div>

      <div class="max-h-96 overflow-y-auto">
        <%= if length(@telemetry_logs) > 0 do %>
          <%= for {log, index} <- Enum.with_index(Enum.take(@telemetry_logs, 10)) do %>
            <div class={"p-4 border-b border-gray-700 hover:bg-gray-750 transition-colors #{if index == 0, do: "bg-blue-900/20"}"}>
              <div class="flex items-center justify-between mb-2">
                <div class="flex items-center space-x-3">
                  <span class={"w-2 h-2 rounded-full #{if index == 0, do: "bg-green-400 animate-pulse", else: "bg-blue-400"}"}>
                  </span>
                  <span class="font-medium text-gray-300">Node: <%= log.node_id %></span>
                  <%= if index == 0 do %>
                    <span class="text-xs bg-green-500 text-white px-2 py-1 rounded">Latest</span>
                  <% end %>
                </div>
                <span class="text-sm text-gray-500">
                  <%= DateTime.from_unix!(log.timestamp_unix) |> DateTime.to_time() |> Time.to_string() %>
                </span>
              </div>

              <div class="grid grid-cols-1 md:grid-cols-3 gap-3 mt-3">
                <%= for metric <- log.data do %>
                  <div class="bg-gray-900 rounded px-3 py-2 border border-gray-600">
                    <div class="text-xs text-gray-400 uppercase tracking-wide">
                      <%= String.replace(metric["name"], "_", " ") %>
                    </div>
                    <div class="text-lg font-semibold text-white flex items-baseline space-x-1">
                      <span><%= :erlang.float_to_binary(metric["value"], decimals: 2) %></span>
                      <span class="text-xs text-gray-400"><%= metric["unit"] %></span>
                    </div>
                  </div>
                <% end %>
              </div>
            </div>
          <% end %>
        <% else %>
          <div class="p-8 text-center text-gray-500">
            <p>No activity logs yet. Send some telemetry data to see it here!</p>
          </div>
        <% end %>
      </div>
    </div>

    <!-- Footer Stats -->
    <div class="mt-8 text-center text-gray-500 text-sm">
      <p>Dashboard last updated: <%= DateTime.utc_now() |> DateTime.to_string() %></p>
    </div>
  </div>
  """
end
end
