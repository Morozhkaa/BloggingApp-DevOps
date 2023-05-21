import React from 'react';
import Post from './Post';

class Posts extends React.Component {
    render() {
        if (this.props.posts.length > 0)
            return (<div>
                {this.props.posts.map((elem) => (
                    <Post onEdit={this.props.onEdit} onDelete={this.props.onDelete} key={elem.post_id} post={elem}/>
                ))}
            </div>)
        else
            return (<div>
                    <h3>No posts yet :( </h3>
            </div>)
    }
}

export default Posts