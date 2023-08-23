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
exports.CopyBlockTypeCopyInstructionsToJSON = exports.CopyBlockTypeCopyInstructionsFromJSONTyped = exports.CopyBlockTypeCopyInstructionsFromJSON = exports.instanceOfCopyBlockTypeCopyInstructions = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the CopyBlockTypeCopyInstructions interface.
 */
function instanceOfCopyBlockTypeCopyInstructions(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCopyBlockTypeCopyInstructions = instanceOfCopyBlockTypeCopyInstructions;
function CopyBlockTypeCopyInstructionsFromJSON(json) {
    return CopyBlockTypeCopyInstructionsFromJSONTyped(json, false);
}
exports.CopyBlockTypeCopyInstructionsFromJSON = CopyBlockTypeCopyInstructionsFromJSON;
function CopyBlockTypeCopyInstructionsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'rooms': !(0, runtime_1.exists)(json, 'rooms') ? undefined : json['rooms'],
        'rateCode': !(0, runtime_1.exists)(json, 'rateCode') ? undefined : json['rateCode'],
        'blockComments': !(0, runtime_1.exists)(json, 'blockComments') ? undefined : json['blockComments'],
        'blockCode': !(0, runtime_1.exists)(json, 'blockCode') ? undefined : json['blockCode'],
        'catering': !(0, runtime_1.exists)(json, 'catering') ? undefined : json['catering'],
        'events': !(0, runtime_1.exists)(json, 'events') ? undefined : json['events'],
        'eventComments': !(0, runtime_1.exists)(json, 'eventComments') ? undefined : json['eventComments'],
        'resources': !(0, runtime_1.exists)(json, 'resources') ? undefined : json['resources'],
        'resourceComments': !(0, runtime_1.exists)(json, 'resourceComments') ? undefined : json['resourceComments'],
        'resourcePrices': !(0, runtime_1.exists)(json, 'resourcePrices') ? undefined : json['resourcePrices'],
        'attendeesCount': !(0, runtime_1.exists)(json, 'attendeesCount') ? undefined : json['attendeesCount'],
        'contractBilling': !(0, runtime_1.exists)(json, 'contractBilling') ? undefined : json['contractBilling'],
        'groupProfile': !(0, runtime_1.exists)(json, 'groupProfile') ? undefined : json['groupProfile'],
        'alternateDates': !(0, runtime_1.exists)(json, 'alternateDates') ? undefined : json['alternateDates'],
        'ratesOfAlternateDates': !(0, runtime_1.exists)(json, 'ratesOfAlternateDates') ? undefined : json['ratesOfAlternateDates'],
        'potentialProfile': !(0, runtime_1.exists)(json, 'potentialProfile') ? undefined : json['potentialProfile'],
        'adjustDecisionAndFollowupDate': !(0, runtime_1.exists)(json, 'adjustDecisionAndFollowupDate') ? undefined : json['adjustDecisionAndFollowupDate'],
        'createAsSubBlock': !(0, runtime_1.exists)(json, 'createAsSubBlock') ? undefined : json['createAsSubBlock'],
        'createAsTourBlock': !(0, runtime_1.exists)(json, 'createAsTourBlock') ? undefined : json['createAsTourBlock'],
        'overbook': !(0, runtime_1.exists)(json, 'overbook') ? undefined : json['overbook'],
        'contractGrid': !(0, runtime_1.exists)(json, 'contractGrid') ? undefined : json['contractGrid'],
        'changeNotes': !(0, runtime_1.exists)(json, 'changeNotes') ? undefined : json['changeNotes'],
        'otherResources': !(0, runtime_1.exists)(json, 'otherResources') ? undefined : json['otherResources'],
        'eventForecastFigures': !(0, runtime_1.exists)(json, 'eventForecastFigures') ? undefined : json['eventForecastFigures'],
    };
}
exports.CopyBlockTypeCopyInstructionsFromJSONTyped = CopyBlockTypeCopyInstructionsFromJSONTyped;
function CopyBlockTypeCopyInstructionsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'rooms': value.rooms,
        'rateCode': value.rateCode,
        'blockComments': value.blockComments,
        'blockCode': value.blockCode,
        'catering': value.catering,
        'events': value.events,
        'eventComments': value.eventComments,
        'resources': value.resources,
        'resourceComments': value.resourceComments,
        'resourcePrices': value.resourcePrices,
        'attendeesCount': value.attendeesCount,
        'contractBilling': value.contractBilling,
        'groupProfile': value.groupProfile,
        'alternateDates': value.alternateDates,
        'ratesOfAlternateDates': value.ratesOfAlternateDates,
        'potentialProfile': value.potentialProfile,
        'adjustDecisionAndFollowupDate': value.adjustDecisionAndFollowupDate,
        'createAsSubBlock': value.createAsSubBlock,
        'createAsTourBlock': value.createAsTourBlock,
        'overbook': value.overbook,
        'contractGrid': value.contractGrid,
        'changeNotes': value.changeNotes,
        'otherResources': value.otherResources,
        'eventForecastFigures': value.eventForecastFigures,
    };
}
exports.CopyBlockTypeCopyInstructionsToJSON = CopyBlockTypeCopyInstructionsToJSON;
