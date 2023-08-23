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
import type { AddressInfoType } from './AddressInfoType';
import {
    AddressInfoTypeFromJSON,
    AddressInfoTypeFromJSONTyped,
    AddressInfoTypeToJSON,
} from './AddressInfoType';
import type { CompanyType } from './CompanyType';
import {
    CompanyTypeFromJSON,
    CompanyTypeFromJSONTyped,
    CompanyTypeToJSON,
} from './CompanyType';
import type { CustomerType } from './CustomerType';
import {
    CustomerTypeFromJSON,
    CustomerTypeFromJSONTyped,
    CustomerTypeToJSON,
} from './CustomerType';
import type { EmailInfoType } from './EmailInfoType';
import {
    EmailInfoTypeFromJSON,
    EmailInfoTypeFromJSONTyped,
    EmailInfoTypeToJSON,
} from './EmailInfoType';
import type { OwnerType } from './OwnerType';
import {
    OwnerTypeFromJSON,
    OwnerTypeFromJSONTyped,
    OwnerTypeToJSON,
} from './OwnerType';
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
import type { TelephoneInfoType } from './TelephoneInfoType';
import {
    TelephoneInfoTypeFromJSON,
    TelephoneInfoTypeFromJSONTyped,
    TelephoneInfoTypeToJSON,
} from './TelephoneInfoType';
import type { URLInfoType } from './URLInfoType';
import {
    URLInfoTypeFromJSON,
    URLInfoTypeFromJSONTyped,
    URLInfoTypeToJSON,
} from './URLInfoType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * 
 * @export
 * @interface RelationshipProfileType
 */
export interface RelationshipProfileType {
    /**
     * 
     * @type {CustomerType}
     * @memberof RelationshipProfileType
     */
    customer?: CustomerType;
    /**
     * 
     * @type {CompanyType}
     * @memberof RelationshipProfileType
     */
    company?: CompanyType;
    /**
     * 
     * @type {TelephoneInfoType}
     * @memberof RelationshipProfileType
     */
    telephone?: TelephoneInfoType;
    /**
     * 
     * @type {AddressInfoType}
     * @memberof RelationshipProfileType
     */
    address?: AddressInfoType;
    /**
     * 
     * @type {EmailInfoType}
     * @memberof RelationshipProfileType
     */
    email?: EmailInfoType;
    /**
     * 
     * @type {URLInfoType}
     * @memberof RelationshipProfileType
     */
    uRLs?: URLInfoType;
    /**
     * 
     * @type {OwnerType}
     * @memberof RelationshipProfileType
     */
    primaryOwner?: OwnerType;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof RelationshipProfileType
     */
    profileIdList?: Array<UniqueIDType>;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof RelationshipProfileType
     */
    changeProfileIdList?: Array<UniqueIDType>;
    /**
     * Indicates if this relationship is the primary relationship.
     * @type {string}
     * @memberof RelationshipProfileType
     */
    primary?: string;
    /**
     * Relationship identifier.
     * @type {string}
     * @memberof RelationshipProfileType
     */
    id?: string;
    /**
     * 
     * @type {ProfileStatusType}
     * @memberof RelationshipProfileType
     */
    statusCode?: ProfileStatusType;
    /**
     * 
     * @type {ProfileTypeType}
     * @memberof RelationshipProfileType
     */
    profileType?: ProfileTypeType;
}

/**
 * Check if a given object implements the RelationshipProfileType interface.
 */
export function instanceOfRelationshipProfileType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RelationshipProfileTypeFromJSON(json: any): RelationshipProfileType {
    return RelationshipProfileTypeFromJSONTyped(json, false);
}

export function RelationshipProfileTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RelationshipProfileType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'customer': !exists(json, 'customer') ? undefined : CustomerTypeFromJSON(json['customer']),
        'company': !exists(json, 'company') ? undefined : CompanyTypeFromJSON(json['company']),
        'telephone': !exists(json, 'telephone') ? undefined : TelephoneInfoTypeFromJSON(json['telephone']),
        'address': !exists(json, 'address') ? undefined : AddressInfoTypeFromJSON(json['address']),
        'email': !exists(json, 'email') ? undefined : EmailInfoTypeFromJSON(json['email']),
        'uRLs': !exists(json, 'uRLs') ? undefined : URLInfoTypeFromJSON(json['uRLs']),
        'primaryOwner': !exists(json, 'primaryOwner') ? undefined : OwnerTypeFromJSON(json['primaryOwner']),
        'profileIdList': !exists(json, 'profileIdList') ? undefined : ((json['profileIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'changeProfileIdList': !exists(json, 'changeProfileIdList') ? undefined : ((json['changeProfileIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'primary': !exists(json, 'primary') ? undefined : json['primary'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'statusCode': !exists(json, 'statusCode') ? undefined : ProfileStatusTypeFromJSON(json['statusCode']),
        'profileType': !exists(json, 'profileType') ? undefined : ProfileTypeTypeFromJSON(json['profileType']),
    };
}

export function RelationshipProfileTypeToJSON(value?: RelationshipProfileType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'customer': CustomerTypeToJSON(value.customer),
        'company': CompanyTypeToJSON(value.company),
        'telephone': TelephoneInfoTypeToJSON(value.telephone),
        'address': AddressInfoTypeToJSON(value.address),
        'email': EmailInfoTypeToJSON(value.email),
        'uRLs': URLInfoTypeToJSON(value.uRLs),
        'primaryOwner': OwnerTypeToJSON(value.primaryOwner),
        'profileIdList': value.profileIdList === undefined ? undefined : ((value.profileIdList as Array<any>).map(UniqueIDTypeToJSON)),
        'changeProfileIdList': value.changeProfileIdList === undefined ? undefined : ((value.changeProfileIdList as Array<any>).map(UniqueIDTypeToJSON)),
        'primary': value.primary,
        'id': value.id,
        'statusCode': ProfileStatusTypeToJSON(value.statusCode),
        'profileType': ProfileTypeTypeToJSON(value.profileType),
    };
}

