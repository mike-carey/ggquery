# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Ggquery < Formula
  desc ""
  homepage ""
  version "1.0.0"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/mike-carey/ggquery/releases/download/v1.0.0/ggquery_1.0.0_darwin_amd64.tar.gz"
      sha256 "10ffe1cd229b1de47a2004f8825a0eda88fc5b8987f1e0cb3d1398bea2bd6eca"

      def install
        bin.install "ggquery"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/mike-carey/ggquery/releases/download/v1.0.0/ggquery_1.0.0_darwin_arm64.tar.gz"
      sha256 "111c606006924f082a29a4607fb43bd055f7b4bcf185875cebf4ff9e9513600c"

      def install
        bin.install "ggquery"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/mike-carey/ggquery/releases/download/v1.0.0/ggquery_1.0.0_linux_arm64.tar.gz"
      sha256 "c5522a20df512068db8fa66fcc2a261426abb7754f621d3ec9513abe4fabacdc"

      def install
        bin.install "ggquery"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/mike-carey/ggquery/releases/download/v1.0.0/ggquery_1.0.0_linux_amd64.tar.gz"
      sha256 "a223890b6fb0e58e1d84d31eeb4af20774b54c34f7dc8f6054d1e4b23294ae39"

      def install
        bin.install "ggquery"
      end
    end
  end
end