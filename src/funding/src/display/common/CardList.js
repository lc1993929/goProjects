import {Card, Icon, Image, List, ListContent, Progress} from "semantic-ui-react";


const CardList = (props) => {
    const details = props.details;
    // console.table(details);

    let cards = details.map(detail => {
        return <CardExampleImageCard key={detail.fundingAddress} detail={detail} onCardClick={props.onCardClick}/>
    })

    return (
        <Card.Group itemsPerRow={4}>
            {cards}
        </Card.Group>
    )
}

export default CardList;


const CardExampleImageCard = (props) => {
    let detail = props.detail;
    let percent = detail.balance / detail.targetMoney * 100;
    let onCardClick = props.onCardClick;

    return (
        <Card onClick={() => {
            onCardClick(detail);
        }}>
            <Image src='logo192.png' wrapped ui={false}/>
            <Card.Content>
                <Card.Header>{detail.projectName}</Card.Header>
                <Card.Meta>
                    <span className='date'>剩余时间：{detail.remainTime}</span>
                    <Progress percent={percent} progress size='small'/>
                </Card.Meta>
            </Card.Content>
            <Card.Content extra>
                <List horizontal style={{display: 'flex', justifyContent: 'space-around'}}>
                    <List.Item>
                        <List.Content>
                            <List.Header>已筹</List.Header>
                            {detail.balance} wei
                        </List.Content>
                    </List.Item>
                    <List.Item>
                        <List.Content>
                            <List.Header>目标</List.Header>
                            {detail.targetMoney}
                        </List.Content>
                    </List.Item>
                    <List.Item>
                        <List.Content>
                            <List.Header>参与人数</List.Header>
                            {detail.supportersCount}
                        </List.Content>
                    </List.Item>
                </List>
            </Card.Content>
        </Card>
    )
}