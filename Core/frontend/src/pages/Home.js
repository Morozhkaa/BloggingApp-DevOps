import React from 'react';
import Posts from '../components/Posts';
import AddPost from '../components/AddPost';
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

//const basePostUrl = "https://myapp.com/api/post-service/v1/"
//const basePostUrl = "http://localhost:9000/api/post-service/v1/"
const basePostUrl = process.env.REACT_APP_postURL

class Post extends React.Component {
    constructor(props) {
        super(props)
        axios.get(basePostUrl+"getAll").then((res) => {
            console.log(res.data)
            this.setState({posts: res.data})
        })

        this.state = {
            posts : [],
        }

        this.addPost = this.addPost.bind(this)
        this.deletePost = this.deletePost.bind(this)
        this.editPost = this.editPost.bind(this)
    }

    render() {
        return (<div>
            <h2>Post list</h2>
            <main>
                <Posts posts={this.state.posts} onEdit={this.editPost} onDelete={this.deletePost}/>
                <ToastContainer /> 
            </main>
            <aside>
                <AddPost onAdd={this.addPost}/>
            </aside>
        </div>)
    }

    // --------------------------------------------------------------------------------------

    async deletePost(id) {
        let statusCode = 0;

        async function status() {
            const response = await axios.delete(basePostUrl+"delete/"+id, {
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
                posts: this.state.posts.filter((el) => el.post_id !== id),
                isForbidden: true
            })
        } else if (statusCode === 403) {
            toast('🤔 У вас недостаточно прав для удаления данного поста!', {
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


    // --------------------------------------------------------------------------------------

    async editPost(post) {
        let data = {};
        let statusCode = 0;
        async function status() {
            const response = await axios.post(basePostUrl+ "update/"+post.post_id, post, {
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
            let allPosts = [];
            for (let item of this.state.posts) {
                if (item.post_id === post.post_id) {
                    allPosts.push(data)
                } else {
                    allPosts.push(item)
                }
            }
            this.setState({posts: []}, () => {
                this.setState({
                    posts: [...allPosts],
                })
            })
        } else if (statusCode === 403) {
            toast('🤔 У вас недостаточно прав для редактирования данного поста!', {
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

    // --------------------------------------------------------------------------------------

    async addPost(post) {
        let data = {};
        let statusCode = 0;
        async function status() {
            const response = await axios.post(basePostUrl+"create", post, {
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
            const id = data.post_id
            post.post_id = id
            post.author.email = data.email
            post.author.login = data.login
            post.created_at.date = data.created_at.date
            post.created_at.time = data.created_at.time
            post.updated_at.date = data.updated_at.date
            post.updated_at.time = data.updated_at.time
            this.setState({
                posts: [...this.state.posts, {id, ...post}],
            })
        } else if (statusCode === 403) {
            toast('🦄 Только авторизованные пользователи могут публиковать посты!', {
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