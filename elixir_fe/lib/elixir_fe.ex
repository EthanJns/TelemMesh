defmodule ElixirFe do
  @moduledoc """
  Documentation for `ElixirFe`.
  """

  @doc """
  Hello world.

  ## Examples
      mix run -e "ElixirFe.hello";
      mix deps.get to install dependencies


  """

  use Application
  alias UUID
  def start(_type, _args) do
    IO.puts("Application Start")
    IO.puts(UUID.uuid4())
    main()
    #The following line starts a supervisor which manages child processes
    #The one_for_one is so if a child process dies, only that dies
    Supervisor.start_link([], strategy: :one_for_one)
  end

  def main() do
    x = 5
    IO.puts(x)
  end
end
