import React, {Component} from "react";
import {getAllFundings} from "../../eth/interaction";
import CardList from "../common/CardList";
import HandleSupportForm from "./HandleSupportForm";

class AllFundingTab extends Component {

    constructor(props) {
        super(props);
        this.state = {
            allFundings: [],
        };

        this.childRef = React.createRef()
    }

    async componentDidMount() {

        let allFundings = await getAllFundings();

        this.setState({
            allFundings: allFundings
        })
    }

    handleChildEvent = (ref) => {
        // 将子组件的实例存到 this.childRef 中, 这样整个父组件就能拿到
        this.childRef = ref
    }

    onCardClick = (detail) => {
        this.childRef.updateSupportForm(detail);
    }


    render() {
        return (
            <div>
                <CardList details={this.state.allFundings} onCardClick={this.onCardClick}/>
                <HandleSupportForm onChildEvent={this.handleChildEvent}/>
            </div>
        )
    }

}

export default AllFundingTab;