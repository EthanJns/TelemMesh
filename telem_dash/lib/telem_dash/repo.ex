defmodule TelemDash.Repo do
  use Ecto.Repo,
    otp_app: :telem_dash,
    adapter: Ecto.Adapters.Postgres
end
