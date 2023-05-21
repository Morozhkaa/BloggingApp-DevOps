import React from 'react';

class AddComment extends React.Component {
    commentAdd = {}
    selectedComm = {}

    constructor(props){
        super(props)
        this.state = {
            comment_id: "",
            post_id: "",
            text: "",
            commenter: {
                email: "",
                login: "",
            },
            created_at: {
                date: "",
                time: "",
            },
            updated_at: {
                date: "",
                time: "",
            }
        }
    }
    render() {
        return (
            <div className='CommentForm'>
                <form ref={(el) => this.myForm =el}>
                    <div className='CommentInput'>
                        <textarea placeholder='Text..' onChange={(e) => this.setState({text: e.target.value})}></textarea>
                        <button type="button" onClick={() => {
                            this.myForm.reset()
                            this.commentAdd = {
                                comment_id: "",
                                post_id: this.props.post_id,
                                text: this.state.text,
                                commenter: {
                                    email: "",
                                    login: "",
                                },
                                created_at: {
                                    date: "",
                                    time: "",
                                },
                                updated_at: {
                                    date: "",
                                    time: "",
                                }
                            }
                            if (this.props.comment)
                                this.commentAdd.comment_id = this.props.comment.comment_id
                            if (this.state.text !== "")
                                this.props.onAddComment(this.commentAdd)
                        }
                        }>Add</button>
                    </div>
                </form>
            </div>
        )
    }
}

export default AddComment