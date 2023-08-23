"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ResvRoutingCriteriaTypeToJSON = exports.ResvRoutingCriteriaTypeFromJSONTyped = exports.ResvRoutingCriteriaTypeFromJSON = exports.instanceOfResvRoutingCriteriaType = void 0;
const runtime_1 = require("../runtime");
const ReservationId_1 = require("./ReservationId");
const ResvRoutingInfoTypeComp_1 = require("./ResvRoutingInfoTypeComp");
const ResvRoutingInfoTypeFolio_1 = require("./ResvRoutingInfoTypeFolio");
const ResvRoutingInfoTypeRequest_1 = require("./ResvRoutingInfoTypeRequest");
const ResvRoutingInfoTypeRoom_1 = require("./ResvRoutingInfoTypeRoom");
/**
 * Check if a given object implements the ResvRoutingCriteriaType interface.
 */
function instanceOfResvRoutingCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfResvRoutingCriteriaType = instanceOfResvRoutingCriteriaType;
function ResvRoutingCriteriaTypeFromJSON(json) {
    return ResvRoutingCriteriaTypeFromJSONTyped(json, false);
}
exports.ResvRoutingCriteriaTypeFromJSON = ResvRoutingCriteriaTypeFromJSON;
function ResvRoutingCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'folio': !(0, runtime_1.exists)(json, 'folio') ? undefined : (0, ResvRoutingInfoTypeFolio_1.ResvRoutingInfoTypeFolioFromJSON)(json['folio']),
        'room': !(0, runtime_1.exists)(json, 'room') ? undefined : (0, ResvRoutingInfoTypeRoom_1.ResvRoutingInfoTypeRoomFromJSON)(json['room']),
        'comp': !(0, runtime_1.exists)(json, 'comp') ? undefined : (0, ResvRoutingInfoTypeComp_1.ResvRoutingInfoTypeCompFromJSON)(json['comp']),
        'request': !(0, runtime_1.exists)(json, 'request') ? undefined : (0, ResvRoutingInfoTypeRequest_1.ResvRoutingInfoTypeRequestFromJSON)(json['request']),
        'refreshFolio': !(0, runtime_1.exists)(json, 'refreshFolio') ? undefined : json['refreshFolio'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !(0, runtime_1.exists)(json, 'reservationId') ? undefined : (0, ReservationId_1.ReservationIdFromJSON)(json['reservationId']),
        'retrievePostingsForRoomRouting': !(0, runtime_1.exists)(json, 'retrievePostingsForRoomRouting') ? undefined : json['retrievePostingsForRoomRouting'],
    };
}
exports.ResvRoutingCriteriaTypeFromJSONTyped = ResvRoutingCriteriaTypeFromJSONTyped;
function ResvRoutingCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'folio': (0, ResvRoutingInfoTypeFolio_1.ResvRoutingInfoTypeFolioToJSON)(value.folio),
        'room': (0, ResvRoutingInfoTypeRoom_1.ResvRoutingInfoTypeRoomToJSON)(value.room),
        'comp': (0, ResvRoutingInfoTypeComp_1.ResvRoutingInfoTypeCompToJSON)(value.comp),
        'request': (0, ResvRoutingInfoTypeRequest_1.ResvRoutingInfoTypeRequestToJSON)(value.request),
        'refreshFolio': value.refreshFolio,
        'hotelId': value.hotelId,
        'reservationId': (0, ReservationId_1.ReservationIdToJSON)(value.reservationId),
        'retrievePostingsForRoomRouting': value.retrievePostingsForRoomRouting,
    };
}
exports.ResvRoutingCriteriaTypeToJSON = ResvRoutingCriteriaTypeToJSON;
