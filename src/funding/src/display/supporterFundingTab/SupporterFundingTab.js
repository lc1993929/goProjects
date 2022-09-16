import {Component} from "react";
import {
    approveRequest,
    getSupporterFundings,
    getSupportersCount,
    listRequestsByFundingAddress
} from "../../eth/interaction";
import CardList from "../common/CardList";
import {Button} from "semantic-ui-react";
import RequestTable from "../common/RequestTable";

class SupporterFundingTab extends Component {


    state = {
        supporterFundings: [],
        chooseFunding: null,
        requests: [],
        supportersCount: 0
    };

    async componentDidMount() {

        let supporterFundings = await getSupporterFundings();

        this.setState({
            supporterFundings: supporterFundings
        })
    }

    onCardClick = (detail) => {
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

    handleApprove = async (index) => {
        await approveRequest(this.state.chooseFunding.fundingAddress, index);
    }


    render() {
        return (
            <div>
                <CardList details={this.state.supporterFundings} onCardClick={this.onCardClick}/>

                {

                    this.state.chooseFunding &&
                    <div>
                        <Button onClick={this.listRequestDetails}>申请花费详情</Button>
                        <RequestTable requests={this.state.requests} handleApprove={this.handleApprove}
                                      supportersCount={this.state.supportersCount} pageKey={1}/>
                    </div>
                }
            </div>
        )
    }

}

export default SupporterFundingTab;