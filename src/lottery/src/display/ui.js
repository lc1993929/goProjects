import React from 'react'
import {Card, Image, Statistic, Button, Icon} from 'semantic-ui-react'

const CardExampleCard = (props) => (
    <Card>
        <Image src='/images/img.png' wrapped ui={false}/>
        <Card.Content>
            <Card.Header>畅的彩票</Card.Header>
            <Card.Meta>
                <p className='date'>管理员：{props.manager}</p>
                <p className='date'>当前用户：{props.currentAccount}</p>
                <p className='date'>上期中奖用户：{props.winner}</p>
            </Card.Meta>
            <Card.Description>
                祝你中奖！！！
            </Card.Description>
        </Card.Content>
        <Card.Content extra>
            <p>{props.playersCount}人参与</p>
        </Card.Content>

        <Card.Content extra>
            <Statistic color='red'>
                <Statistic.Value>{props.balance} ETH</Statistic.Value>
                <Statistic.Label>奖金池</Statistic.Label>
            </Statistic>
        </Card.Content>
        <Card.Content extra>
            <Statistic color='blue'>
                <Statistic.Value>第{props.round}期</Statistic.Value>
                <a href='#'>查看历史开奖</a>
            </Statistic>
        </Card.Content>

        <Button animated='fade' color={"red"} onClick={props.play} disabled={props.isClicked}>
            <Button.Content visible>购买放飞梦想</Button.Content>
            <Button.Content hidden>购买0.000001ETH</Button.Content>
        </Button>
        <Button inverted color='blue' style={{display: props.isShowButton}} onClick={props.openReward}
                disabled={props.isClicked}>
            开奖
        </Button>
        <Button inverted color='orange' style={{display: props.isShowButton}} onClick={props.backMoney}
                disabled={props.isClicked}>
            退奖
        </Button>
    </Card>
)

export default CardExampleCard