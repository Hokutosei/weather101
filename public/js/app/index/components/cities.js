// utility for debugger
var log = function(str) { console.log(str); };

var CityList = React.createClass({
	render: function() {

		return (
            <li>
                { this.props.data.Name } { this.props.data.Sum }
                <Chart data={ this.props.data }/>
            </li>
        )
	}
});

var Chart = React.createClass({
    renderChart: function() {
        console.log(this.props)
        console.log(this.refs)
        console.log("debug---")
        var node = this.refs.chartNode.getDOMNode()
            , dataSeries = this.props.data.Items;

        //console.log(d3)

        var chartOptions = {

            rangeSelector: {
                selected: 4
            },

            yAxis: {
                labels: {
                    formatter: function () {
                        return (this.value > 0 ? ' + ' : '') + this.value + '%';
                    }
                },
                plotLines: [{
                    value: 0,
                    width: 2,
                    color: 'silver'
                }]
            },

            plotOptions: {
                series: {
                    compare: 'percent'
                }
            },

            tooltip: {
                valueDecimals: 2
            },
            chart: {
                renderTo: node ,
                width: 400 ,
                height: 400
            },
            series: dataSeries
        };


        jQuery(function($) {
            console.log("chart jquery initialized");
            log(chartOptions);
            $(node).highcharts('StockChart', chartOptions)
        });
//        var chartInstance = new Highcharts.Chart(chartOptions)
//        this.setState({
//            chartInstance: chartInstance
//        })





    },

    componentWillReceiveProps: function(nextProps) {

    },

    shouldComponentUpdate: function(nextProps, nextState) {
        return nextProps.data.Items.length > 0;
    },

    componentDidUpdate: function(nextProps) {
        this.renderChart();
    },

    render: function() {
        return React.DOM.div({className: "chart", ref: "chartNode" })
    }
});

var Cities = React.createClass({
	getInitialState: function() {
		return {
			Data: []
		}
	},

	componentWillReceiveProps: function(nextProps) {
		this.setState({ Data: nextProps })
	},

	render: function() {
		var city_data = [{ Name: "Loading.." }];
		if(this.props.Data) {
			city_data = this.props.Data
		}		
		return (
			<ul>
				{
					city_data.map(function(city) {
						return <CityList data={ city } />;
					})
				}					

			</ul>				
		)
	}
});

app.value('Cities', Cities);