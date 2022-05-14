# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Marketxiv < Formula
  desc ""
  homepage ""
  version "0.9.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/BitsOfAByte/MarketXIV/releases/download/v0.9.0/marketxiv_darwin_amd64.zip"
      sha256 "683db88e9a51154dea3c441f377322d229c5e165097e5af4942616fc59f81501"

      def install
        bin.install "marketxiv"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/BitsOfAByte/MarketXIV/releases/download/v0.9.0/marketxiv_darwin_arm64.zip"
      sha256 "fb3ea81f54b15c364ee648195949c0e6370df56b4de10e959afce51b71b74274"

      def install
        bin.install "marketxiv"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/BitsOfAByte/MarketXIV/releases/download/v0.9.0/marketxiv_linux_amd64.zip"
      sha256 "7c28f73498682747d129252fa94afc4d74f76054e80ed214600bb25572e75c14"

      def install
        bin.install "marketxiv"
      end
    end
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/MarketXIV/releases/download/v0.9.0/marketxiv_linux_arm.zip"
      sha256 "2b8439776c64085d4ee3dd9c93c34cccae8d2cb8254dde39e174b9814fe96ef0"

      def install
        bin.install "marketxiv"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/MarketXIV/releases/download/v0.9.0/marketxiv_linux_arm64.zip"
      sha256 "2da93b353ac10505c9b9ee9a02cf4a0515a0393c1bd590794b522e73d3aff36e"

      def install
        bin.install "marketxiv"
      end
    end
  end
end
