import React, { Component } from 'react';
import './App.css';

class App extends Component {
  constructor() {
    super();
    this.state = {
      title: "",
      description: "",
      author: "",
      cover_image: "",
    };
  }

  componentDidMount() {
    fetch('http://localhost:5000')
      .then(response => response.json())
      .then(results => {
        console.log(results)
        this.setState(
        { 
          title: results[0].title,
          description: results[0].description,
          author: results[0].author,
          cover_image: results[0].cover_image,
        })
      })
      .catch(err => console.error(this.props.url, err.toString()))
  }

  render() {
    return (
      <div className="outer-div">
        <img src={this.state.cover_image} alt="unavailable" className="book-cover"/>
        <h2 className="title">{this.state.title}</h2>
        <h2 className="author">{this.state.author}</h2>
        <p className="description">{this.state.description}</p>
      </div>
    )
  }
}

export default App;
