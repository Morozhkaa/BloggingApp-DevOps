import React from 'react';
import {AiFillCloseCircle} from "react-icons/ai";
import {AiTwotoneEdit, AiOutlineComment} from "react-icons/ai";
import Comments from './Comments'
import AddComment from './AddComment';
import AddPost from './AddPost';
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

//const baseCommUrl = "https://myapp.com/api/comms-service/v1/"
//const baseCommUrl = "http://localhost:9090/api/comms-service/v1/"
const baseCommUrl = process.env.REACT_APP_commentURL

class Post extends React.Component {
    constructor(props) {
        super(props)

        axios.get(baseCommUrl+"getAll/"+this.props.post.post_id).then((res) => {
            res.data != null && this.setState({comments: res.data})
        })
        this.state = {
            editForm: false,
            commentForm: false,
            comments: []
        }
        this.addComment = this.addComment.bind(this)
        this.editComment = this.editComment.bind(this)
        this.deleteComment = this.deleteComment.bind(this)
    }

    post = this.props.post
    render() {
        return (
            <div className='Post'>
                    <AiFillCloseCircle onClick={() => this.props.onDelete(this.post.post_id)} className='delete-icon' />
                    <AiTwotoneEdit className='edit-icon' onClick={() => {
                        this.setState({
                            editForm: !this.state.editForm
                        })
                    }}/>
                    <h3 className='PostOutput'>{this.post.title}</h3>
                    <p className='PostOutput'><b>Description: </b> "{this.post.description}"</p>
                    <p className='PostOutput'><b>Content: </b></p>
                    <p className='PostOutputContent'>{this.post.content}</p>
                    <p className='PostOutput'> <b>Author: </b>{this.post.author.login}</p>
                    <p className='PostOutput'><b>Email: </b>{this.post.author.email} </p>
                    <p className='PostOutput'><b>Created: </b>{this.post.created_at.date}  {this.post.created_at.time}</p>
                    <p className='PostOutput'><b>Updated: </b>{this.post.updated_at.date}  {this.post.updated_at.time}</p>
                    <AiOutlineComment className='comment-icon' onClick={() => {
                        this.setState({
                            commentForm: !this.state.commentForm
                        })
                    }}/>
                    
                    {this.state.editForm && <AddPost post={this.post} onAdd={this.props.onEdit} />}
                    <div className='CommList'> {this.state.commentForm && <Comments comments={this.state.comments}
                         onDeleteComment={this.deleteComment} onEditComment={this.editComment} post_id={this.post.post_id}/>}
                    </div>
                    <aside> {this.state.commentForm && <AddComment onAddComment={this.addComment} post_id={this.post.post_id}/>} </aside>
            </div>
        )
    }

    async addComment(comment) {
        <ToastContainer />
        let data = {};
        let statusCode = 0;
        async function status() {
            const response = await axios.post(baseCommUrl+"create/"+comment.post_id, comment, {
                dataType: 'json',
                withCredentials: true,
              }).catch(function (error) {
                statusCode = error.response.status
                return
              })
            if (statusCode === 0) {
                data = response.data;
                statusCode = response.status
            }
        }
        await status();
        console.log(data);
        console.log(statusCode);

        if (statusCode === 201) {
            const id = data.comm_id
            comment.comment_id = id
            comment.commenter.email = data.email
            comment.commenter.login = data.login
            comment.created_at.date = data.created_at.date
            comment.created_at.time = data.created_at.time
            comment.updated_at.date = data.updated_at.date
            comment.updated_at.time = data.updated_at.time
            this.setState({
                comments: [...this.state.comments, {id, ...comment}]
            })
        } else if (statusCode === 403) {
            toast('🦄 Только авторизованные пользователи могут оставлять комментарии!', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        } else if (statusCode >= 500) {
            toast('😅 Возникли неполадки: один из сервисов auth/comms не доступен', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        }
    }

    // --------------------------------------------------------------------------------------

    async editComment(comment) {
        let data = {};
        let statusCode = 0;
        async function status() {
            const response = await axios.post(baseCommUrl+"update/"+comment.comment_id, comment, {
                dataType: 'json',
                withCredentials: true,
              }).catch(function (error) {
                statusCode = error.response.status
                return
              })
            if (statusCode === 0) {
                data = response.data;
                statusCode = response.status
            }
        }
        await status();
        console.log(data);
        console.log(statusCode);

        if (statusCode === 200) {
            let allComms = [];
            for (let item of this.state.comments) {
                if (item.comment_id === comment.comment_id) {
                    allComms.push(data)
                } else {
                    allComms.push(item)
                }
            }
            this.setState({comments: []}, () => {
                this.setState({comments: [...allComms]})
            })
        } else if (statusCode === 403) {
            toast('🤔 У вас недостаточно прав для редактирования данного комментария!', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        } else if (statusCode >= 500) {
            toast('😅 Возникли неполадки: один из сервисов auth/comms не доступен', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        }
    }

    // --------------------------------------------------------------------------------------

    async deleteComment(comm_id) {
        let statusCode = 0;

        async function status() {
            const response = await axios.delete(baseCommUrl+"delete/"+comm_id, {
                withCredentials: true,
              }).catch(function (error) {
                statusCode = error.response.status
                return
              })
            if (statusCode === 0) {
                statusCode = response.status
            }
        }
        await status();
        console.log(statusCode);

        if (statusCode === 200) {
            this.setState({
                comments: this.state.comments.filter((el) => el.comment_id !== comm_id)
            })
        } else if (statusCode === 403) {
            toast('🤔 У вас недостаточно прав для удаления данного комментария!', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        } else if (statusCode >= 500) {
            toast('😅 Возникли неполадки: один из сервисов auth/post не доступен', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        }
    }
}

export default Post