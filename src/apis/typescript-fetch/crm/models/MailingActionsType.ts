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
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';

/**
 * Defines mailing action list for the profile.
 * @export
 * @interface MailingActionsType
 */
export interface MailingActionsType {
    /**
     * Defines mailing action code and description.
     * @type {Array<CodeDescriptionType>}
     * @memberof MailingActionsType
     */
    mailingAction?: Array<CodeDescriptionType>;
    /**
     * When true indicates that profile has subscribed to the mailing list.
     * @type {boolean}
     * @memberof MailingActionsType
     */
    active?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof MailingActionsType
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof MailingActionsType
     */
    count?: number;
}

/**
 * Check if a given object implements the MailingActionsType interface.
 */
export function instanceOfMailingActionsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MailingActionsTypeFromJSON(json: any): MailingActionsType {
    return MailingActionsTypeFromJSONTyped(json, false);
}

export function MailingActionsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MailingActionsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'mailingAction': !exists(json, 'mailingAction') ? undefined : ((json['mailingAction'] as Array<any>).map(CodeDescriptionTypeFromJSON)),
        'active': !exists(json, 'active') ? undefined : json['active'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !exists(json, 'count') ? undefined : json['count'],
    };
}

export function MailingActionsTypeToJSON(value?: MailingActionsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'mailingAction': value.mailingAction === undefined ? undefined : ((value.mailingAction as Array<any>).map(CodeDescriptionTypeToJSON)),
        'active': value.active,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}

