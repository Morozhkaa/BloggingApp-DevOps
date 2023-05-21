import React from 'react';
import Comment from './Comment';

class Comments extends React.Component {
    render() {
        if (this.props.comments.length > 0) {
            return (<div>
                {this.props.comments.map((elem) => (
                    <Comment onEditComment={this.props.onEditComment} onDeleteComment={this.props.onDeleteComment} key={elem.comment_id} comment={elem} post_id={this.props.post_id}/>
                ))}
            </div>)
        } else
            return (<div className='comment'>
                <p>No comments yet..</p>
            </div>)
        
    }
}

export default Comments