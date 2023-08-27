# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Httpurr < Formula
  desc ""
  homepage "https://github.com/rednafi/httpurr"
  version "0.1.0-alpha"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/rednafi/httpurr/rednafi/httpurr/releases/download/v0.1.0-alpha/httpurr_Darwin_arm64.tar.gz"
      sha256 "6a64bc6bc69e694604eac818ca2129c15cd92e5b46cdf5cb3142daee614250d7"

      def install
        bin.install "httpurr"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/rednafi/httpurr/rednafi/httpurr/releases/download/v0.1.0-alpha/httpurr_Darwin_x86_64.tar.gz"
      sha256 "00f6a1bde5db0fb06413bf483d2801872ce569bb500b33d450fc4dd38911ba45"

      def install
        bin.install "httpurr"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/rednafi/httpurr/rednafi/httpurr/releases/download/v0.1.0-alpha/httpurr_Linux_arm64.tar.gz"
      sha256 "522515cb3c10ee959aac32d7f7321067eb4a82e50a0874a8756c51b24161d16b"

      def install
        bin.install "httpurr"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/rednafi/httpurr/rednafi/httpurr/releases/download/v0.1.0-alpha/httpurr_Linux_x86_64.tar.gz"
      sha256 "fe1d86fb7ca8bcb1168f036396bb757c3cec714b7e3357b4327742ecf73ce41a"

      def install
        bin.install "httpurr"
      end
    end
  end
end
