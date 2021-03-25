import './App.css';
import { Component } from "react"
import { Button, Form, Loading, Select } from 'element-react';


class App extends Component {

  state = {
    form: {},
    types: {},
    loading: true,
    services: []
  }

  async componentDidMount() {
    // retrieve services from remote
    const { services } = await (await fetch("/services")).json()
    this.setState({ loading: false, services })
  }

  render() {
    return (
      <div className="App">
        <Loading loading={this.state.loading}>

          <h1 className="title">
            Archive Life
          </h1>

          <Form
            ref="form"
            model={this.state.form}
            labelWidth="100"
          >

            <Form.Item label="Archive Type">
              <Select value={this.state.form.type} placeholder="the archive type">
                {
                  this.state.services.map(
                    service => <Select.Option key={service.Name} label={service.Name} ></Select.Option>
                  )
                }
              </Select>
            </Form.Item>

            <Form.Item>
              <Button type="info">Execute</Button>
            </Form.Item>

          </Form>
        </Loading>
      </div>
    )
  }
}

export default App;
