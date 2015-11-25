class TopController < ApplicationController
  def index
    @messages = Message.all
  end
end
