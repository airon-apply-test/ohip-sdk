"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.EffectiveRateTypeToJSON = exports.EffectiveRateTypeFromJSONTyped = exports.EffectiveRateTypeFromJSON = exports.instanceOfEffectiveRateType = void 0;
const runtime_1 = require("../runtime");
const RateByAgeBucketType_1 = require("./RateByAgeBucketType");
/**
 * Check if a given object implements the EffectiveRateType interface.
 */
function instanceOfEffectiveRateType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfEffectiveRateType = instanceOfEffectiveRateType;
function EffectiveRateTypeFromJSON(json) {
    return EffectiveRateTypeFromJSONTyped(json, false);
}
exports.EffectiveRateTypeFromJSON = EffectiveRateTypeFromJSON;
function EffectiveRateTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'onePersonRate': !(0, runtime_1.exists)(json, 'onePersonRate') ? undefined : json['onePersonRate'],
        'twoPersonRate': !(0, runtime_1.exists)(json, 'twoPersonRate') ? undefined : json['twoPersonRate'],
        'threePersonRate': !(0, runtime_1.exists)(json, 'threePersonRate') ? undefined : json['threePersonRate'],
        'fourPersonRate': !(0, runtime_1.exists)(json, 'fourPersonRate') ? undefined : json['fourPersonRate'],
        'fivePersonRate': !(0, runtime_1.exists)(json, 'fivePersonRate') ? undefined : json['fivePersonRate'],
        'extraPersonRate': !(0, runtime_1.exists)(json, 'extraPersonRate') ? undefined : json['extraPersonRate'],
        'extraChildRate': !(0, runtime_1.exists)(json, 'extraChildRate') ? undefined : json['extraChildRate'],
        'oneChildRate': !(0, runtime_1.exists)(json, 'oneChildRate') ? undefined : json['oneChildRate'],
        'twoChildrenRate': !(0, runtime_1.exists)(json, 'twoChildrenRate') ? undefined : json['twoChildrenRate'],
        'threeChildrenRate': !(0, runtime_1.exists)(json, 'threeChildrenRate') ? undefined : json['threeChildrenRate'],
        'fourChildrenRate': !(0, runtime_1.exists)(json, 'fourChildrenRate') ? undefined : json['fourChildrenRate'],
        'rateByAgeBuckets': !(0, runtime_1.exists)(json, 'rateByAgeBuckets') ? undefined : (json['rateByAgeBuckets'].map(RateByAgeBucketType_1.RateByAgeBucketTypeFromJSON)),
        'minimumChildrenForFreeStay': !(0, runtime_1.exists)(json, 'minimumChildrenForFreeStay') ? undefined : json['minimumChildrenForFreeStay'],
        'pointsRequired': !(0, runtime_1.exists)(json, 'pointsRequired') ? undefined : json['pointsRequired'],
        'overrideFloorAmount': !(0, runtime_1.exists)(json, 'overrideFloorAmount') ? undefined : json['overrideFloorAmount'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'ratePlanCode': !(0, runtime_1.exists)(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : json['roomType'],
        'numberOfRooms': !(0, runtime_1.exists)(json, 'numberOfRooms') ? undefined : json['numberOfRooms'],
        'amountBeforeTax': !(0, runtime_1.exists)(json, 'amountBeforeTax') ? undefined : json['amountBeforeTax'],
        'start': !(0, runtime_1.exists)(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !(0, runtime_1.exists)(json, 'end') ? undefined : (new Date(json['end'])),
    };
}
exports.EffectiveRateTypeFromJSONTyped = EffectiveRateTypeFromJSONTyped;
function EffectiveRateTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'onePersonRate': value.onePersonRate,
        'twoPersonRate': value.twoPersonRate,
        'threePersonRate': value.threePersonRate,
        'fourPersonRate': value.fourPersonRate,
        'fivePersonRate': value.fivePersonRate,
        'extraPersonRate': value.extraPersonRate,
        'extraChildRate': value.extraChildRate,
        'oneChildRate': value.oneChildRate,
        'twoChildrenRate': value.twoChildrenRate,
        'threeChildrenRate': value.threeChildrenRate,
        'fourChildrenRate': value.fourChildrenRate,
        'rateByAgeBuckets': value.rateByAgeBuckets === undefined ? undefined : (value.rateByAgeBuckets.map(RateByAgeBucketType_1.RateByAgeBucketTypeToJSON)),
        'minimumChildrenForFreeStay': value.minimumChildrenForFreeStay,
        'pointsRequired': value.pointsRequired,
        'overrideFloorAmount': value.overrideFloorAmount,
        'hotelId': value.hotelId,
        'ratePlanCode': value.ratePlanCode,
        'roomType': value.roomType,
        'numberOfRooms': value.numberOfRooms,
        'amountBeforeTax': value.amountBeforeTax,
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0, 10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0, 10)),
    };
}
exports.EffectiveRateTypeToJSON = EffectiveRateTypeToJSON;
