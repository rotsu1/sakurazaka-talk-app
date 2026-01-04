//
//  Utils.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import Foundation

let formatterSimple = {
    let formatter = DateFormatter()
    formatter.dateFormat = "yyyy/MM/dd"
    return formatter
}()

let formatterDetailed = {
    let formatter = DateFormatter()
    formatter.dateFormat = "yyyy/MM/dd HH:mm"
    return formatter
}()