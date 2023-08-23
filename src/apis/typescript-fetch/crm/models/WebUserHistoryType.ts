/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Creation date time, Creator Id, last modification date time and last Modifier Id.
 * @export
 * @interface WebUserHistoryType
 */
export interface WebUserHistoryType {
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof WebUserHistoryType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof WebUserHistoryType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof WebUserHistoryType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof WebUserHistoryType
     */
    lastModifierId?: string;
    /**
     * Last login date.
     * @type {Date}
     * @memberof WebUserHistoryType
     */
    lastLogin?: Date;
    /**
     * Last Password change date.
     * @type {Date}
     * @memberof WebUserHistoryType
     */
    lastPasswordChange?: Date;
    /**
     * Inactive date.
     * @type {Date}
     * @memberof WebUserHistoryType
     */
    inactiveDate?: Date;
}

/**
 * Check if a given object implements the WebUserHistoryType interface.
 */
export function instanceOfWebUserHistoryType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function WebUserHistoryTypeFromJSON(json: any): WebUserHistoryType {
    return WebUserHistoryTypeFromJSONTyped(json, false);
}

export function WebUserHistoryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): WebUserHistoryType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'createDateTime': !exists(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !exists(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !exists(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !exists(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'lastLogin': !exists(json, 'lastLogin') ? undefined : (new Date(json['lastLogin'])),
        'lastPasswordChange': !exists(json, 'lastPasswordChange') ? undefined : (new Date(json['lastPasswordChange'])),
        'inactiveDate': !exists(json, 'inactiveDate') ? undefined : (new Date(json['inactiveDate'])),
    };
}

export function WebUserHistoryTypeToJSON(value?: WebUserHistoryType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'lastLogin': value.lastLogin === undefined ? undefined : (value.lastLogin.toISOString().substr(0,10)),
        'lastPasswordChange': value.lastPasswordChange === undefined ? undefined : (value.lastPasswordChange.toISOString().substr(0,10)),
        'inactiveDate': value.inactiveDate === undefined ? undefined : (value.inactiveDate.toISOString().substr(0,10)),
    };
}

