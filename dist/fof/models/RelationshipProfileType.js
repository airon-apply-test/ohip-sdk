"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RelationshipProfileTypeToJSON = exports.RelationshipProfileTypeFromJSONTyped = exports.RelationshipProfileTypeFromJSON = exports.instanceOfRelationshipProfileType = void 0;
const runtime_1 = require("../runtime");
const AddressInfoType_1 = require("./AddressInfoType");
const CompanyType_1 = require("./CompanyType");
const CustomerType_1 = require("./CustomerType");
const EmailInfoType_1 = require("./EmailInfoType");
const OwnerType_1 = require("./OwnerType");
const ProfileStatusType_1 = require("./ProfileStatusType");
const ProfileTypeType_1 = require("./ProfileTypeType");
const TelephoneInfoType_1 = require("./TelephoneInfoType");
const URLInfoType_1 = require("./URLInfoType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the RelationshipProfileType interface.
 */
function instanceOfRelationshipProfileType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRelationshipProfileType = instanceOfRelationshipProfileType;
function RelationshipProfileTypeFromJSON(json) {
    return RelationshipProfileTypeFromJSONTyped(json, false);
}
exports.RelationshipProfileTypeFromJSON = RelationshipProfileTypeFromJSON;
function RelationshipProfileTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'customer': !(0, runtime_1.exists)(json, 'customer') ? undefined : (0, CustomerType_1.CustomerTypeFromJSON)(json['customer']),
        'company': !(0, runtime_1.exists)(json, 'company') ? undefined : (0, CompanyType_1.CompanyTypeFromJSON)(json['company']),
        'telephone': !(0, runtime_1.exists)(json, 'telephone') ? undefined : (0, TelephoneInfoType_1.TelephoneInfoTypeFromJSON)(json['telephone']),
        'address': !(0, runtime_1.exists)(json, 'address') ? undefined : (0, AddressInfoType_1.AddressInfoTypeFromJSON)(json['address']),
        'email': !(0, runtime_1.exists)(json, 'email') ? undefined : (0, EmailInfoType_1.EmailInfoTypeFromJSON)(json['email']),
        'uRLs': !(0, runtime_1.exists)(json, 'uRLs') ? undefined : (0, URLInfoType_1.URLInfoTypeFromJSON)(json['uRLs']),
        'primaryOwner': !(0, runtime_1.exists)(json, 'primaryOwner') ? undefined : (0, OwnerType_1.OwnerTypeFromJSON)(json['primaryOwner']),
        'profileIdList': !(0, runtime_1.exists)(json, 'profileIdList') ? undefined : (json['profileIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'changeProfileIdList': !(0, runtime_1.exists)(json, 'changeProfileIdList') ? undefined : (json['changeProfileIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'primary': !(0, runtime_1.exists)(json, 'primary') ? undefined : json['primary'],
        'id': !(0, runtime_1.exists)(json, 'id') ? undefined : json['id'],
        'statusCode': !(0, runtime_1.exists)(json, 'statusCode') ? undefined : (0, ProfileStatusType_1.ProfileStatusTypeFromJSON)(json['statusCode']),
        'profileType': !(0, runtime_1.exists)(json, 'profileType') ? undefined : (0, ProfileTypeType_1.ProfileTypeTypeFromJSON)(json['profileType']),
    };
}
exports.RelationshipProfileTypeFromJSONTyped = RelationshipProfileTypeFromJSONTyped;
function RelationshipProfileTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'customer': (0, CustomerType_1.CustomerTypeToJSON)(value.customer),
        'company': (0, CompanyType_1.CompanyTypeToJSON)(value.company),
        'telephone': (0, TelephoneInfoType_1.TelephoneInfoTypeToJSON)(value.telephone),
        'address': (0, AddressInfoType_1.AddressInfoTypeToJSON)(value.address),
        'email': (0, EmailInfoType_1.EmailInfoTypeToJSON)(value.email),
        'uRLs': (0, URLInfoType_1.URLInfoTypeToJSON)(value.uRLs),
        'primaryOwner': (0, OwnerType_1.OwnerTypeToJSON)(value.primaryOwner),
        'profileIdList': value.profileIdList === undefined ? undefined : (value.profileIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'changeProfileIdList': value.changeProfileIdList === undefined ? undefined : (value.changeProfileIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'primary': value.primary,
        'id': value.id,
        'statusCode': (0, ProfileStatusType_1.ProfileStatusTypeToJSON)(value.statusCode),
        'profileType': (0, ProfileTypeType_1.ProfileTypeTypeToJSON)(value.profileType),
    };
}
exports.RelationshipProfileTypeToJSON = RelationshipProfileTypeToJSON;
