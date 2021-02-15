import React from 'react'
import { Chart } from 'react-charts'
import { Client } from '@stomp/stompjs';
import WhisperClient from '../../app/WhisperClient'

function LineGraph() {

    // const client = new Client({
    //     brokerURL: 'ws://localhost:8081/ws-crypto-monitor-broadcaster/websocket',
    //     debug: function (str) {
    //         console.log(str);
    //     },
    //     reconnectDelay: 5000,
    //     heartbeatIncoming: 4000,
    //     heartbeatOutgoing: 4000,
    // });

    // client.onConnect = function (frame) {
    //     this.subscribe('/topic/tracker', function (message) {
    //         console.log(message)
    //     });
    //     this.publish({destination: "/app/update", body: "TEST"})
    // };
    //
    // client.onUnhandledMessage = function (message) {
    //     console.log(message)
    // }
    //
    // client.onStompError = function (frame) {
    //     // Will be invoked in case of error encountered at Broker
    //     // Bad login/passcode typically will cause an error
    //     // Complaint brokers will set `message` header with a brief message. Body may contain details.
    //     // Compliant brokers will terminate the connection after any error
    //     console.log('Broker reported error: ' + frame.headers['message']);
    //     console.log('Additional details: ' + frame.body);
    // };
    //
    // client.activate();

    // const data = [
    //     {
    //         label: 'Sentiment',
    //         data: [[0, 1], [1, 2], [2, 4], [3, 2], [4, 10]]
    //     },
    //     {
    //         label: 'Price',
    //         data: [[0, 3], [1, 1], [2, 5], [3, 6], [4, 4]]
    //     }
    // ]

    // React.useMemo(
    //     () => Client.get('/price-sentiment?assetpairs=XBTUSD')
    //         .then(res => {
    //             //console.log(res);
    //             // data[0].data.push([8, 8])
    //             // data[1].data.push([1, 1])
    //             // console.log(data[0]);
    //             // console.log(data[1]);
    //         }),
    //     []
    // )

    // const axes = React.useMemo(
    //     () => [
    //         { primary: true, type: 'linear', position: 'bottom' },
    //         { type: 'linear', position: 'left' }
    //     ],
    //     []
    // )
    //
    // React.useEffect(() => {
    //     // const assetPairs = Client.get('/price-sentiment?assetpairs=XBTUSD')
    //     //     .then(res => {
    //     //         console.log(res);
    //     //         console.log(data[0].data.push([8, 8]));
    //     //         console.log(data[1].data.push([1, 1]));
    //     //     });
    // });

    WhisperClient.get('/')
        .then((response) => {
            console.log("RESPONSE")
            console.log(response.data);
        });

    const lineChart = (
        // A react-chart hyper-responsively and continuously fills the available
        // space of its parent element automatically
        // <Chart data={data} axes={axes} className="line-graph-main"/>
        <div></div>
    )

    return lineChart;
}

export default LineGraph;