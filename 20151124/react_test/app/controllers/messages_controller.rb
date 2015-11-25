class MessagesController < ApplicationController
  def index
    messages = Message.all
    render json: messages
  end

  def create
    message = Message.new(create_params)
    if message.save
      render json: message, status: :created
    else
      render json: message, status: :unprocessable
    end
  end

  private

  def create_params
    params.permit(:user, :text)
  end
end