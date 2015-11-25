var MessageBox = React.createClass({
  getInitialState: function() {
    return { messages: this.props.messages };
  },

  handleMessageSubmit: function(message) {
    $.ajax({
      url: this.props.url,
      dataType: 'json',
      type: 'POST',
      data: message,
      success: function(message) {
        console.log(message);
        var newMessages = this.state.messages.concat(message);
        this.setState({ messages: newMessages });
      }.bind(this),
      error: function(_xhr, status, err) {
        console.log(this.props.url,status,err.toString());
      }.bind(this)
    });
  },

  render: function() {
    var messageItems = this.state.messages.map(function(message){
      return (
        <MessageItem key={message.id} message={message} />
      );
    });

    return (
      <div className="messageBox">
        {messageItems}
        <MessageForm onMessageSubmit={this.handleMessageSubmit} />
      </div>
    )
  }
});