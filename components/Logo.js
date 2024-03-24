import React from 'react'
import { View, Image, Dimensions } from 'react-native'

export default function Logo(props) {
    const img = require('../assets/ScriptBuyLogo.png')
    const width = Dimensions.get('window').width * props.ratio
    const resizeRatio = width / Image.resolveAssetSource(img).width 
    
    return (
        <View className="items-center pb-3 h-fit">
            <Image
                className={'p-' + props.padding}
                style={{
                    width: width,
                    height: resizeRatio * Image.resolveAssetSource(img).height
                }}
                resizeMode={props.resizeMode}
                source={img}
            />
        </View>
    )
}

Logo.defaultProps = {
    ratio: 0.5,
    padding: 4,
    resizeMode: 'contain'
}