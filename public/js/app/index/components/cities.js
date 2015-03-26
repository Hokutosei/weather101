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
        console.log("debug---")
        var node = this.refs.chartNode.getDomNode()
            , dataSeries = this.props.data.Items;

        jQuery(function($) {
            $(node).highcharts({
                chart: {
                    plotBackgroundColor: "#EFEFEF",
                    height: 300,
                    type: 'bar'
                },
                series: dataSeries
            })
        })
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