import React, { Component } from 'react';
import './ChatInput.scss';

class ChatInput extends Component {
  
  render() {
    return (
      <div className='ChatInput'>
        <input onKeyDown={this.props.send} placeholder="Type a message... Press enter to send"  maxLength={300}/>
      </div>
    );
  };

}

export default ChatInput;
