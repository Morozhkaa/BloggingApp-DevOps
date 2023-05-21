import React from 'react';

class AddPost extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            post_id: "",
            title: "",
            description: "",
            content: "",
            author: {
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
            <div className='PostForm'>
            <form ref={(el) => this.myForm = el}>
                <input className='PostInput' placeholder='Title' onChange={(e) => this.setState({title: e.target.value})}/>
                <input className='PostInput' placeholder='Description' onChange={(e) => this.setState({description: e.target.value})}/>
                <textarea className='PostInput' placeholder='Content' onChange={(e) => this.setState({content: e.target.value})}></textarea>
                <button type ='button' onClick={() => {
                    this.myForm.reset()
                    this.postAdd = {
                        post_id: "",
                        title: this.state.title,
                        description: this.state.description,
                        content: this.state.content,
                        author: {
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
                    if (this.props.post) {
                        this.postAdd.post_id = this.props.post.post_id
                        this.postAdd.author.email = this.props.post.author.email
                        this.postAdd.author.login = this.props.post.author.login
                    }
                    this.props.onAdd(this.postAdd)
                }
                }>Publish</button>
            </form>
            </div>
        )
    }
}

export default AddPost