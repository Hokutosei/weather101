var CityInput = React.createClass({
    render: function() {
        console.log(this.props)
        return (
            <div className="input-group">
              <div className="input-group-addon">{ this.props.index }</div>
              <input type="text" className="form-control" id="exampleInputAmount" placeholder={ this.props.item } />
            </div>
        )
    }
})

var CityList = React.createClass({
    getInitialState: function() {
        return {
            cityListArr: []
        }
    },

    render: function() {
        console.log(this.props)
        return (
            <div>
                <ul className="city_list_input">
                    {   _.map(this.props.city_list, function(item, index) {
                            return (
                                <li>
                                    <CityInput item={ item } index={ index } />
                                </li>
                            )
                        })
                    }
                </ul>

            </div>
        )
    }

});
app.value('CityList', CityList)
