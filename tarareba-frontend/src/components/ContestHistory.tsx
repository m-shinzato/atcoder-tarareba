import React from 'react'
import { GET_CONTEST_HISTORY } from '../graphql/tags/getContestHistory'

import { useQuery } from 'react-apollo-hooks'

import { Table, Card, Button, Checkbox } from 'semantic-ui-react'

import getRatingColorStyle from '../utils/getRatingColorStyle'

interface Contest {
  isRated: boolean
  place: number
  actualOldRating: number
  actualNewRating: number
  performance: number
  innerPerformance: number
  contestScreenName: string
  contestName: string
  contestNameEn: string
  endTime: string
  optimalOldRating: number
  optimalNewRating: number
  isParticipated: boolean
}

interface Contests {
  contestsByUserID: Contest[]
}

type Props = {
  userID: string
}

const ContestHistory: React.FC<Props> = (props) => {
  // マウント時にリクエストが走る
  const { loading, error, data } = useQuery<Contests>(GET_CONTEST_HISTORY, {
    variables: {
      userID: props.userID,
    },
  })

  if (loading) return <div>loading</div>
  if (error) {
    return (
      <div>
        <p>なんらかのエラーが発生しました。</p>
        <p>
          <a href="https://twitter.com/monkukui2">こちら</a>にご連絡ください
        </p>
      </div>
    )
  }

  if (data!.contestsByUserID.length == 0) {
    console.log(data)
    if (props.userID != '') {
      return (
        <div>
          <p>ユーザーが見つかりません</p>
        </div>
      )
    }

    return <></>
  }

  const currentRating = data!.contestsByUserID[
    data!.contestsByUserID.length - 1
  ].actualNewRating
  const optimalRating = data!.contestsByUserID[
    data!.contestsByUserID.length - 1
  ].optimalNewRating

  return (
    <>
      <div style={{ marginTop: '5em' }}>
        <Card>
          <Card.Content>
            <Card.Header>結果</Card.Header>
          </Card.Content>
          <Table definition style={{ width: '90%', margin: '0 auto' }}>
            <Table.Body>
              <Table.Row>
                <Table.Cell>実際のレート</Table.Cell>
                <Table.Cell style={getRatingColorStyle(currentRating)}>
                  {currentRating}
                </Table.Cell>
              </Table.Row>
              <Table.Row>
                <Table.Cell>架空のレート</Table.Cell>
                <Table.Cell style={getRatingColorStyle(optimalRating)}>
                  {optimalRating}
                </Table.Cell>
              </Table.Row>
              <Table.Row>
                <Table.Cell>差分</Table.Cell>
                <Table.Cell style={{ fontWeight: 'bold' }}>
                  +{optimalRating - currentRating}
                </Table.Cell>
              </Table.Row>
            </Table.Body>
          </Table>
          <Card.Content style={{ marginTop: '2em' }}>
            <Button basic color="green">
              Twitter に投稿
            </Button>
          </Card.Content>
        </Card>
        <Table celled={true} style={{ fontSize: 'calc(6px + 1vmin)' }}>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>日付け</Table.HeaderCell>
              <Table.HeaderCell>コンテスト</Table.HeaderCell>
              <Table.HeaderCell>順位</Table.HeaderCell>
              <Table.HeaderCell>パフォーマンス</Table.HeaderCell>
              <Table.HeaderCell>実際の Rating</Table.HeaderCell>
              <Table.HeaderCell>架空の Rating</Table.HeaderCell>
              <Table.HeaderCell>参加・不参加</Table.HeaderCell>
            </Table.Row>
          </Table.Header>

          <Table.Body>
            {data!.contestsByUserID.map((record) => {
              return (
                <Table.Row key={record.endTime}>
                  <Table.Cell>{record.endTime}</Table.Cell>
                  <Table.Cell>
                    <a
                      target="_blank"
                      href={'https://' + record.contestScreenName}
                    >
                      {record.contestName}
                    </a>
                  </Table.Cell>
                  <Table.Cell>
                    <a
                      target="_blank"
                      href={
                        'https://' +
                        record.contestScreenName +
                        '/standings?watching=' +
                        props.userID
                      }
                    >
                      {record.place}
                    </a>
                  </Table.Cell>
                  {record.isRated ? (
                    <>
                      <Table.Cell
                        style={getRatingColorStyle(record.performance)}
                      >
                        {record.performance}
                      </Table.Cell>
                      <Table.Cell
                        style={getRatingColorStyle(record.actualNewRating)}
                      >
                        {record.actualNewRating}
                        &nbsp;
                        {record.actualNewRating > record.actualOldRating ? (
                          <>
                            (+{record.actualNewRating - record.actualOldRating})
                          </>
                        ) : (
                          <>
                            (-{record.actualOldRating - record.actualNewRating})
                          </>
                        )}
                      </Table.Cell>
                      {record.isParticipated ? (
                        <>
                          <Table.Cell
                            style={getRatingColorStyle(record.optimalNewRating)}
                          >
                            {record.optimalNewRating}
                            &nbsp;
                            {record.optimalNewRating >
                            record.optimalOldRating ? (
                              <>
                                (+
                                {record.optimalNewRating -
                                  record.optimalOldRating}
                                )
                              </>
                            ) : (
                              <>
                                (-
                                {record.optimalOldRating -
                                  record.optimalNewRating}
                                )
                              </>
                            )}
                          </Table.Cell>
                        </>
                      ) : (
                        <Table.Cell>-</Table.Cell>
                      )}
                      <Table.Cell>
                        <Checkbox />
                      </Table.Cell>
                    </>
                  ) : (
                    <>
                      <Table.Cell>-</Table.Cell>
                      <Table.Cell>-</Table.Cell>
                      <Table.Cell>-</Table.Cell>
                      <Table.Cell>-</Table.Cell>
                    </>
                  )}
                </Table.Row>
              )
            })}
          </Table.Body>
        </Table>
      </div>
    </>
  )
}

export default ContestHistory
