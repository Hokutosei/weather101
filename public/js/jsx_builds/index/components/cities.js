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



var CityLatestTemp = React.createClass({displayName: "CityLatestTemp",

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

	weatherDescription: function(items) {
		var lastItem = _.last(items)
		if(lastItem == undefined) return false;

		return lastItem.description
	},

	render: function() {

		var lastestRecord = this.latestItem(this.state.Items)
			, tempDescription = this.weatherDescription(this.state.Items)

		return (
			React.createElement("div", null, 
				React.createElement("span", {className: "cityTemp", style:  this.state.style},  lastestRecord ), " ", React.createElement("br", null), 
				React.createElement("span", null, " ",  tempDescription, " ")
			)
		)
	}
});

var CityListItem = React.createClass({displayName: "CityListItem",
	render: function() {

		var br_style = { clear: 'both' }
		log(this.props)
		log("key debug ------")

		return (
            React.createElement("li", {className: "city", key:  this.props.data.id}, 
				React.createElement("div", {className: ""}, 
					React.createElement("div", {className: "col-sm-7"}, 
		                 this.props.data.Name, " Chart", 
		                React.createElement(Chart, {chart_data:  this.props.data})
					), 
					React.createElement("div", {className: "col-sm-2 col-xs-offset-2"}, 
						React.createElement("div", {className: "row"}, 
							React.createElement("div", {className: "cityNameProfile"}, 
								 this.props.data.Name
							), 

							React.createElement("div", {className: "cityDataTotalRecords"}, 
								 this.props.data.Sum, " records"
							), 

							React.createElement("div", {className: "cityLatestTemp"}, 
								React.createElement(CityLatestTemp, {Items:  this.props.data.Items})
							)
						)
					)
				), 
				React.createElement("br", {style:  br_style })
            )
        )
	}
});

var Chart = React.createClass({displayName: "Chart",
    renderChart: function(node) {
            var dataSeries = this.props.chart_data.Items
				, chartName = this.props.chart_data.Name

				if(dataSeries == 0) return false;

				var node_data = []

				async.each(dataSeries, function(item, callback) {

					var tempVal = !item['celsius'] || item['celsius'] == 0 ? convertCelsius(item['temp']) : item['celsius']
					node_data.push([dateFormatter(item['created_at']), tempVal])
					callback()
				}, function(err) {
					log(chartName)
					log("--- debug")
					log(node_data.length)
					chartInit(node_data)
				})



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
						navigator : {
			                adaptToUpdatedData: false,
							enabled: false,
			                series : {
			                    data : node_data
			                }
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
								dataGrouping: {
									enabled: true
								},
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

	componentDidUpdate: function() {
	},

    shouldComponentUpdate: function(nextProps, nextState) {
		log(nextProps.chart_data.Items.length)
		log("debug shouldComponentUpdate-----")
        return nextProps.chart_data.Items != undefined && nextProps.chart_data.Items.length > 0;
    },

    componentDidUpdate: function(nextProps) {
		var node = this.refs.chartNode.getDOMNode()
        this.renderChart(node);
    },

    render: function() {
        return React.DOM.div({className: "chart", ref: "chartNode" })
    }
});

var Cities = React.createClass({displayName: "Cities",
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
			React.createElement("ul", {className: "cities"}, 
				
					city_data.map(function(city, index) {
						console.log(index)
						city['id'] = index;
						return React.createElement(CityListItem, {data:  city, key:  index });
					})
				
			)
		)
	}
});

app.value('Cities', Cities);
