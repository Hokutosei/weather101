// utility for debugger
var log = function(str) { console.log(str); };

function dateFormatter(date_str) {
	var date = new Date(date_str.slice(0, date_str.indexOf(".")))
	return date.getTime()
}

var CityList = React.createClass({
	render: function() {

		return (
            <li>
                { this.props.data.Name } { this.props.data.Sum }
                <Chart chart_data={ this.props.data }/>
            </li>
        )
	}
});

var Chart = React.createClass({
    renderChart: function() {
        var node = this.refs.chartNode.getDOMNode()
            , dataSeries = this.props.chart_data.Items
						, chartName = this.props.chart_data.Name

				jQuery(function($) {
					var node_data = []
					for(var i = 0; i < 500; i++) {
						var ds = dataSeries[i]
						// node_data.push([ds['temp'], ds['created_at'].slice(0, 19)])
						node_data.push([dateFormatter(ds['created_at']), ds['temp']])
						// node_data.push(ds['temp'])
					}
					log(node_data)

					$(node).highcharts('StockChart', {
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
                  pointFormat: '<span style="color:{series.color}">{series.name}</span>: <b>{point.y}</b> ({point.change}%)<br/>',
                  valueDecimals: 2
              },

              series: {
								name: chartName,
								data: dataSeries
							}
          });

				})
    },

    componentWillReceiveProps: function(nextProps) {

    },

    shouldComponentUpdate: function(nextProps, nextState) {
        return nextProps.chart_data.Items.length > 0;
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
