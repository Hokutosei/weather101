// utility for debugger
var log = function(str) { console.log(str); };

function dateFormatter(date_str) {
	var date = new Date(date_str.slice(0, date_str.indexOf(".")))
	return date.getTime()
}

function convertCelsius(temp) {
	return parseInt((temp - 273.15).toFixed(2))
}

function tempColor(temp) {
	var css = {};
	switch (true) {
		case (temp >= 1 && temp <= 5):
			css.color = '#bedff6'
			break;
		case (temp >= 6 && temp <= 8):
			css.color = '#8ac6ef'
			break;
		case (temp >= 29 && temp <= 33):
			css.color = '#ffb732'
			break;
		default: {
			css.color = '#000'
		}
	}
	return css
}



var CityLatestTemp = React.createClass({

	getInitialState: function() {
		return {
			Items: [],
			style: {}
		}
	},

	componentWillReceiveProps: function(nextProps) {
		this.setState({ Items: nextProps.Items })
	},

	shouldComponentUpdate: function(nextProps, nextState) {
		return nextProps.Items != undefined && nextProps.Items.length > 0;
	},

	componentDidUpdate: function(nextProps) {
	},

	latestItem: function(items) {
		var lastItem = _.last(items);
		if(lastItem == undefined) return false;

		var convertedTemp = convertCelsius(lastItem.temp)
		this.state.style = tempColor(parseInt(convertedTemp))

		return convertedTemp + '\u00B0' + 'C'
	},


	render: function() {

		var lastestRecord = this.latestItem(this.state.Items)

		return (
			<span className="cityTemp" style={ this.state.style }>{ lastestRecord }</span>
		)
	}
});

var CityListItem = React.createClass({
	render: function() {

		return (
            <li className="city">
				<div className="">
					<div className="col-sm-7">
		                { this.props.data.Name } Chart
		                <Chart chart_data={ this.props.data }/>
					</div>
					<div className="col-sm-2 col-xs-offset-2">
						<div className="row">
							<div className="cityNameProfile">
								{ this.props.data.Name }
							</div>

							<div className="cityDataTotalRecords">
								{ this.props.data.Sum } records
							</div>

							<div className="cityLatestTemp">
								<CityLatestTemp Items={ this.props.data.Items } />
							</div>
						</div>
					</div>
				</div>
            </li>
        )
	}
});

var Chart = React.createClass({
    renderChart: function(node) {
            var dataSeries = this.props.chart_data.Items
				, chartName = this.props.chart_data.Name
				, pointLimit = dataSeries.length

				var node_data = []
				for(var i = 0; i < pointLimit; i++) {
					var ds = dataSeries[i]
					// node_data.push([ds['temp'], ds['created_at'].slice(0, 19)])
					//node_data.push([dateFormatter(ds['created_at']), ds['temp']])
					var tempVal = !ds['celsius'] || ds['celsius'] == 0 ? convertCelsius(ds['temp']) : ds['celsius']
					node_data.push([dateFormatter(ds['created_at']), tempVal])
					if(node_data.length == pointLimit) {
						chartInit(node_data)
						return false
					}
				};

				function chartInit(node_data) {
					var startDate = (new Date(dataSeries[0].created_at))

					$(node).highcharts('StockChart', {
						chart: {
							zoomType: 'x',
							width: 800,
							height: 200,
							panning: false
						},
						yAxis: {
							title: {
								text: 'temp'
							}
						},
						navigator: {
						            enabled: false
						},
						scrollbar: {
							enabled: false
						},
						rangeSelector : {
						        enabled: false
						},
						series: [
							{
								name: chartName,
								data: node_data,
								pointStart: (new Date(dataSeries[0].created_at).getTime()),
								pointInterval: 24 * 3200
							}
						]
					});

				}
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
			<ul className="cities">
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
