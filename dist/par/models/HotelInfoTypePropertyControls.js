"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.HotelInfoTypePropertyControlsToJSON = exports.HotelInfoTypePropertyControlsFromJSONTyped = exports.HotelInfoTypePropertyControlsFromJSON = exports.instanceOfHotelInfoTypePropertyControls = void 0;
const runtime_1 = require("../runtime");
const HotelInfoTypePropertyControlsApplicationMode_1 = require("./HotelInfoTypePropertyControlsApplicationMode");
const HotelInfoTypePropertyControlsCateringCurrencyFormatting_1 = require("./HotelInfoTypePropertyControlsCateringCurrencyFormatting");
const HotelInfoTypePropertyControlsCurrencyFormatting_1 = require("./HotelInfoTypePropertyControlsCurrencyFormatting");
const HotelInfoTypePropertyControlsDateTimeFormatting_1 = require("./HotelInfoTypePropertyControlsDateTimeFormatting");
const HotelInfoTypePropertyControlsSellControls_1 = require("./HotelInfoTypePropertyControlsSellControls");
/**
 * Check if a given object implements the HotelInfoTypePropertyControls interface.
 */
function instanceOfHotelInfoTypePropertyControls(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelInfoTypePropertyControls = instanceOfHotelInfoTypePropertyControls;
function HotelInfoTypePropertyControlsFromJSON(json) {
    return HotelInfoTypePropertyControlsFromJSONTyped(json, false);
}
exports.HotelInfoTypePropertyControlsFromJSON = HotelInfoTypePropertyControlsFromJSON;
function HotelInfoTypePropertyControlsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'sellControls': !(0, runtime_1.exists)(json, 'sellControls') ? undefined : (0, HotelInfoTypePropertyControlsSellControls_1.HotelInfoTypePropertyControlsSellControlsFromJSON)(json['sellControls']),
        'currencyFormatting': !(0, runtime_1.exists)(json, 'currencyFormatting') ? undefined : (0, HotelInfoTypePropertyControlsCurrencyFormatting_1.HotelInfoTypePropertyControlsCurrencyFormattingFromJSON)(json['currencyFormatting']),
        'cateringCurrencyFormatting': !(0, runtime_1.exists)(json, 'cateringCurrencyFormatting') ? undefined : (0, HotelInfoTypePropertyControlsCateringCurrencyFormatting_1.HotelInfoTypePropertyControlsCateringCurrencyFormattingFromJSON)(json['cateringCurrencyFormatting']),
        'dateTimeFormatting': !(0, runtime_1.exists)(json, 'dateTimeFormatting') ? undefined : (0, HotelInfoTypePropertyControlsDateTimeFormatting_1.HotelInfoTypePropertyControlsDateTimeFormattingFromJSON)(json['dateTimeFormatting']),
        'applicationMode': !(0, runtime_1.exists)(json, 'applicationMode') ? undefined : (0, HotelInfoTypePropertyControlsApplicationMode_1.HotelInfoTypePropertyControlsApplicationModeFromJSON)(json['applicationMode']),
    };
}
exports.HotelInfoTypePropertyControlsFromJSONTyped = HotelInfoTypePropertyControlsFromJSONTyped;
function HotelInfoTypePropertyControlsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'sellControls': (0, HotelInfoTypePropertyControlsSellControls_1.HotelInfoTypePropertyControlsSellControlsToJSON)(value.sellControls),
        'currencyFormatting': (0, HotelInfoTypePropertyControlsCurrencyFormatting_1.HotelInfoTypePropertyControlsCurrencyFormattingToJSON)(value.currencyFormatting),
        'cateringCurrencyFormatting': (0, HotelInfoTypePropertyControlsCateringCurrencyFormatting_1.HotelInfoTypePropertyControlsCateringCurrencyFormattingToJSON)(value.cateringCurrencyFormatting),
        'dateTimeFormatting': (0, HotelInfoTypePropertyControlsDateTimeFormatting_1.HotelInfoTypePropertyControlsDateTimeFormattingToJSON)(value.dateTimeFormatting),
        'applicationMode': (0, HotelInfoTypePropertyControlsApplicationMode_1.HotelInfoTypePropertyControlsApplicationModeToJSON)(value.applicationMode),
    };
}
exports.HotelInfoTypePropertyControlsToJSON = HotelInfoTypePropertyControlsToJSON;
