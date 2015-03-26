var CityList = React.createClass({
	render: function() {
		return <li>{ this.props.data.Name } { this.props.data.Sum }</li>;
	}
})

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
})

app.value('Cities', Cities)