defmodule ElixirFeTest do
  use ExUnit.Case
  doctest ElixirFe

  test "greets the world" do
    assert ElixirFe.hello() == :world
  end
end
