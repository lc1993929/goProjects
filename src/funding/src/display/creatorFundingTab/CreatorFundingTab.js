import {Component} from "react";
import {
    approveRequest, finalizeRequest,
    getCreatorFundings,
    getSupportersCount,
    listRequestsByFundingAddress
} from "../../eth/interaction";
import CardList from "../common/CardList";
import CreateFundingForm from "./CreateFundingForm";
import HandleRequestForm from "./HandleRequestForm";
import {Button} from "semantic-ui-react";
import RequestTable from "../common/RequestTable";

class CreatorFundingTab extends Component {

    state = {
        creatorFundings: [],
        chooseFunding: null,
        requests: [],
        supportersCount: 0
    };

    async componentDidMount() {

        let creatorFundings = await getCreatorFundings();

        this.setState({
            creatorFundings: creatorFundings
        })
    }


    handleChildEvent = (ref) => {
        // 将子组件的实例存到 this.childRef 中, 这样整个父组件就能拿到
        this.childRef = ref
    }
    onCardClick = (detail) => {
        this.childRef.updateRequestForm(detail);
        this.setState({
            chooseFunding: detail
        })
    }

    listRequestDetails = async () => {
        if (this.state.chooseFunding == null) {
            alert('请先选择一个众筹项目');
            return;
        }
        let requests = await listRequestsByFundingAddress(this.state.chooseFunding.fundingAddress);
        let supportersCount = await getSupportersCount(this.state.chooseFunding.fundingAddress);
        this.setState({requests: requests, supportersCount: supportersCount});

    }

    handleFinalize = async (index) => {
        await finalizeRequest(this.state.chooseFunding.fundingAddress, index);
    }


    render() {
        return (
            <div>
                <CardList details={this.state.creatorFundings} onCardClick={this.onCardClick}/>
                {
                    this.state.chooseFunding &&
                    <div>
                        <Button onClick={this.listRequestDetails}>申请花费详情</Button>
                        <RequestTable requests={this.state.requests} supportersCount={this.state.supportersCount}
                                      handleFinalize={this.handleFinalize} pageKey={2}/>
                    </div>
                }
                <HandleRequestForm onChildEvent={this.handleChildEvent}/>
                <CreateFundingForm/>
            </div>
        )
    }

}

export default CreatorFundingTab;