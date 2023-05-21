import React from 'react';
import {AiOutlineEdit, AiTwotoneDelete} from 'react-icons/ai'
import AddComment from './AddComment';

class Comment extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            editForm: false
        }
    }
    comment = this.props.comment
    render() {
        return (
            <div className='Comment'>
                <AiTwotoneDelete onClick={() => this.props.onDeleteComment(this.comment.comment_id)} className='delete-comment-icon'/>
                <AiOutlineEdit onClick={() => {
                    this.setState({
                        editForm: !this.state.editForm
                    })
                }} className='edit-comment-icon' />
                <p className='Info' style={{color: "rgb(82, 38, 118)"}}><b style={{margin: "5px"}}>{this.comment.commenter.login}</b> {this.comment.created_at.date} {this.comment.created_at.time} <b style={{margin: "10px"}}>upd: {this.comment.updated_at.date} {this.comment.updated_at.time} </b> </p>
                <p style={{margin: "5px"}}>{this.comment.text}</p>

                {this.state.editForm && <AddComment comment={this.comment} onAddComment={this.props.onEditComment} post_id={this.props.post_id}/> }
            </div>
        )
    }
}

export default Comment