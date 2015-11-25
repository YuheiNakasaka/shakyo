Rails.application.routes.draw do
  resources :messages, only: [:index,:create], format: 'json'
  root 'top#index'
end
