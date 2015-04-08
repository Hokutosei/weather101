var CityInput = React.createClass({displayName: "CityInput",
    render: function() {
        console.log(this.props)
        return (
            React.createElement("div", {className: "input-group"}, 
              React.createElement("div", {className: "input-group-addon"},  this.props.index), 
              React.createElement("input", {type: "text", className: "form-control", id: "exampleInputAmount", placeholder:  this.props.item})
            )
        )
    }
})

var CityList = React.createClass({displayName: "CityList",
    getInitialState: function() {
        return {
            cityListArr: []
        }
    },

    render: function() {
        console.log(this.props)
        return (
            React.createElement("div", null, 
                React.createElement("ul", {className: "city_list_input"}, 
                       _.map(this.props.city_list, function(item, index) {
                            return (
                                React.createElement("li", null, 
                                    React.createElement(CityInput, {item:  item, index:  index })
                                )
                            )
                        })
                    
                )

            )
        )
    }

});
app.value('CityList', CityList)
