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

import { exists, mapValues } from '../runtime';
import type { ProfileStatusType } from './ProfileStatusType';
import {
    ProfileStatusTypeFromJSON,
    ProfileStatusTypeFromJSONTyped,
    ProfileStatusTypeToJSON,
} from './ProfileStatusType';
import type { ProfileTypeType } from './ProfileTypeType';
import {
    ProfileTypeTypeFromJSON,
    ProfileTypeTypeFromJSONTyped,
    ProfileTypeTypeToJSON,
} from './ProfileTypeType';
import type { RelationshipAddressType } from './RelationshipAddressType';
import {
    RelationshipAddressTypeFromJSON,
    RelationshipAddressTypeFromJSONTyped,
    RelationshipAddressTypeToJSON,
} from './RelationshipAddressType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * 
 * @export
 * @interface RelationshipProfileSummaryType
 */
export interface RelationshipProfileSummaryType {
    /**
     * Name of the customer
     * @type {string}
     * @memberof RelationshipProfileSummaryType
     */
    customerName?: string;
    /**
     * Name of the company.
     * @type {string}
     * @memberof RelationshipProfileSummaryType
     */
    companyName?: string;
    /**
     * Telephone number assigned to a single location
     * @type {string}
     * @memberof RelationshipProfileSummaryType
     */
    telephoneNumber?: string;
    /**
     * 
     * @type {RelationshipAddressType}
     * @memberof RelationshipProfileSummaryType
     */
    address?: RelationshipAddressType;
    /**
     * Defines the e-mail address.
     * @type {string}
     * @memberof RelationshipProfileSummaryType
     */
    emailAddress?: string;
    /**
     * Unique Code to identify the owner.
     * @type {string}
     * @memberof RelationshipProfileSummaryType
     */
    ownerCode?: string;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof RelationshipProfileSummaryType
     */
    profileIdList?: Array<UniqueIDType>;
    /**
     * Indicates if this relationship is the primary relationship.
     * @type {boolean}
     * @memberof RelationshipProfileSummaryType
     */
    primary?: boolean;
    /**
     * 
     * @type {ProfileStatusType}
     * @memberof RelationshipProfileSummaryType
     */
    profileStatus?: ProfileStatusType;
    /**
     * When true, this is a primary owner.
     * @type {boolean}
     * @memberof RelationshipProfileSummaryType
     */
    primaryOwnerCode?: boolean;
    /**
     * 
     * @type {ProfileTypeType}
     * @memberof RelationshipProfileSummaryType
     */
    profileType?: ProfileTypeType;
}

/**
 * Check if a given object implements the RelationshipProfileSummaryType interface.
 */
export function instanceOfRelationshipProfileSummaryType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RelationshipProfileSummaryTypeFromJSON(json: any): RelationshipProfileSummaryType {
    return RelationshipProfileSummaryTypeFromJSONTyped(json, false);
}

export function RelationshipProfileSummaryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RelationshipProfileSummaryType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'customerName': !exists(json, 'customerName') ? undefined : json['customerName'],
        'companyName': !exists(json, 'companyName') ? undefined : json['companyName'],
        'telephoneNumber': !exists(json, 'telephoneNumber') ? undefined : json['telephoneNumber'],
        'address': !exists(json, 'address') ? undefined : RelationshipAddressTypeFromJSON(json['address']),
        'emailAddress': !exists(json, 'emailAddress') ? undefined : json['emailAddress'],
        'ownerCode': !exists(json, 'ownerCode') ? undefined : json['ownerCode'],
        'profileIdList': !exists(json, 'profileIdList') ? undefined : ((json['profileIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'primary': !exists(json, 'primary') ? undefined : json['primary'],
        'profileStatus': !exists(json, 'profileStatus') ? undefined : ProfileStatusTypeFromJSON(json['profileStatus']),
        'primaryOwnerCode': !exists(json, 'primaryOwnerCode') ? undefined : json['primaryOwnerCode'],
        'profileType': !exists(json, 'profileType') ? undefined : ProfileTypeTypeFromJSON(json['profileType']),
    };
}

export function RelationshipProfileSummaryTypeToJSON(value?: RelationshipProfileSummaryType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'customerName': value.customerName,
        'companyName': value.companyName,
        'telephoneNumber': value.telephoneNumber,
        'address': RelationshipAddressTypeToJSON(value.address),
        'emailAddress': value.emailAddress,
        'ownerCode': value.ownerCode,
        'profileIdList': value.profileIdList === undefined ? undefined : ((value.profileIdList as Array<any>).map(UniqueIDTypeToJSON)),
        'primary': value.primary,
        'profileStatus': ProfileStatusTypeToJSON(value.profileStatus),
        'primaryOwnerCode': value.primaryOwnerCode,
        'profileType': ProfileTypeTypeToJSON(value.profileType),
    };
}

