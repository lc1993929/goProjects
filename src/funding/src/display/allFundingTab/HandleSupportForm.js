import {Dimmer, Form, Label, Loader, Segment} from "semantic-ui-react";
import {Component} from "react";
import {support} from "../../eth/interaction";

class HandleSupportForm extends Component {

    state = {
        active: false,
        projectName: '',
        fundingAddress: '',
        supportMoney: '',
    }

    //子组件完成挂载时, 将子组件的方法 this 作为参数传到父组件的函数中
    componentDidMount() {
        // 在子组件中调用父组件的方法,并把当前的实例传进去
        this.props.onChildEvent(this)
    }

    updateSupportForm = (detail) => {
        this.setState({
            projectName: detail.projectName,
            fundingAddress: detail.fundingAddress,
            supportMoney: detail.supportMoney,
        })
    }

    handleSupport = async () => {
        let {fundingAddress, supportMoney} = this.state
        this.setState({active: true})

        try {
            await support(fundingAddress, supportMoney);
            alert(`支持众筹成功!\n`)
        } catch (e) {
            alert(`支持众筹失败!\n`)
            console.log(e)
        } finally {
            this.setState({active: false})
            window.location.reload()
        }
    }

    render() {
        return (
            <div>
                <h3>参与众筹</h3>
                <Dimmer.Dimmable as={Segment} dimmed={this.state.active}>
                    <Dimmer active={this.state.active} inverted>
                        <Loader>支持中</Loader>
                    </Dimmer>
                    <Form onSubmit={this.handleSupport}>
                        <Form.Input type='text' value={this.state.projectName || ''} label='项目名称:'/>
                        <Form.Input type='text' value={this.state.fundingAddress || ''} label='项目地址:'/>
                        <Form.Input type='text' value={this.state.supportMoney || ''} label='支持金额:'
                                    labelPosition='left'>
                            <Label basic>￥</Label>
                            <input/>
                        </Form.Input>

                        <Form.Button primary content='参与众筹'/>
                    </Form>
                </Dimmer.Dimmable>
            </div>
        )
    }

}

export default HandleSupportForm;