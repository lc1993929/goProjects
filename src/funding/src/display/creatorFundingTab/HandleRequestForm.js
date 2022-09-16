import {Dimmer, Form, Label, Loader, Segment} from "semantic-ui-react";
import React, {Component} from "react";
import {createRequest} from "../../eth/interaction";

class HandleRequestForm extends Component {

    state = {
        active: false,
        projectName: '',
        fundingAddress: '',
        purpose: '',
        cost: 0,
        shopAddress: '',

    }
    handleChange = (e, {name, value}) => this.setState({[name]: value})

    //子组件完成挂载时, 将子组件的方法 this 作为参数传到父组件的函数中
    componentDidMount() {
        // 在子组件中调用父组件的方法,并把当前的实例传进去
        this.props.onChildEvent(this)
    }

    updateRequestForm = (detail) => {
        this.setState({
            projectName: detail.projectName,
            fundingAddress: detail.fundingAddress,
        })
    }

    handleCreateRequest = async () => {
        let {purpose, cost, shopAddress, fundingAddress} = this.state
        this.setState({active: true})

        try {
            await createRequest(purpose, cost, shopAddress, fundingAddress);
            alert(`发起支付请求成功!\n`)
        } catch (e) {
            alert(`发起支付请求失败!\n`)
            console.log(e)
        } finally {
            this.setState({active: false})
            window.location.reload()
        }
    }

    render() {
        let {projectName, fundingAddress, purpose, cost, shopAddress} = this.state;

        return (
            <div>
                <h3>发起付款请求</h3>
                <Dimmer.Dimmable as={Segment} dimmed={this.state.active}>
                    <Dimmer active={this.state.active} inverted>
                        <Loader>请求中</Loader>
                    </Dimmer>
                    <Form onSubmit={this.handleCreateRequest}>
                        <Form.Input type='text' value={projectName || ''} label='项目名称:'/>
                        <Form.Input type='text' value={fundingAddress || ''} label='项目地址:'/>
                        <Form.Input required type='text' placeholder='支付目的' name='purpose'
                                    value={purpose} label='支付目的:' labelPosition='left'
                                    onChange={this.handleChange}>
                            <input/>
                        </Form.Input>
                        <Form.Input required type='text' placeholder='商家地址' name='shopAddress'
                                    value={shopAddress} label='商家地址:' labelPosition='left'
                                    onChange={this.handleChange}>
                            <input/>
                        </Form.Input>
                        <Form.Input required type='text' placeholder='花费金额' name='cost'
                                    value={cost} label='花费金额:' labelPosition='left'
                                    onChange={this.handleChange}>
                            <Label basic>￥</Label>
                            <input/>
                        </Form.Input>

                        <Form.Button primary content='发起申请'/>
                    </Form>
                </Dimmer.Dimmable>
            </div>
        )
    }

}

export default HandleRequestForm;