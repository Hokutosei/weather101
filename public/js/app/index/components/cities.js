// utility for debugger
var log = function(str) { console.log(str); };

function dateFormatter(date_str) {
	var date = new Date(date_str.slice(0, date_str.indexOf(".")))
	return date.getTime()
}

var CityListItem = React.createClass({
	render: function() {

		return (
            <li key={ this.props.key }>
                { this.props.data.Name } { this.props.data.Sum }
                <Chart chart_data={ this.props.data }/>
            </li>
        )
	}
});

var Chart = React.createClass({
    renderChart: function(node) {
            var dataSeries = this.props.chart_data.Items
				, chartName = this.props.chart_data.Name

		jQuery(function($) {
			var node_data = []
			for(var i = 0; i < 500; i++) {
				var ds = dataSeries[i]
				// node_data.push([ds['temp'], ds['created_at'].slice(0, 19)])
				//node_data.push([dateFormatter(ds['created_at']), ds['temp']])
				node_data.push(ds['celsius'])
				if(node_data.length == 500) {
					chartInit()
				}
			};
			function hideZoomBar(chart) {
			        chart.rangeSelector.zoomText.hide();
			        $.each(chart.rangeSelector.buttons, function () {
			            this.hide();
			        });
			        $(chart.rangeSelector.divRelative).hide();
			};
			function chartInit() {
				var startDate = (new Date(dataSeries[0].created_at))
				log(startDate)

				$(node).highcharts('StockChart', {
					chart: {
						zoomType: 'x'
					},
					yAxis: {
						title: {
							text: 'temp'
						}
					},
					navigator: {
					            enabled: false
					},
					series: [
						{
							name: chartName,
							data: node_data,
							pointStart: (new Date(dataSeries[0].created_at).getTime()),
							pointInterval: 3600 * 100
						}
					]
				});

			}
		})

    },

    componentWillReceiveProps: function(nextProps) {

    },

    shouldComponentUpdate: function(nextProps, nextState) {
        return nextProps.chart_data.Items.length > 0;
    },

    componentDidUpdate: function(nextProps) {
		var node = this.refs.chartNode.getDOMNode()
        this.renderChart(node);
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
					city_data.map(function(city, index) {
						return <CityListItem data={ city } key={ index } />;
					})
				}
			</ul>
		)
	}
});

app.value('Cities', Cities);
