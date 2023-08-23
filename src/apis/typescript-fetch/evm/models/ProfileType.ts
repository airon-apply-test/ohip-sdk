/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
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
import type { ImageSetType } from './ImageSetType';
import {
    ImageSetTypeFromJSON,
    ImageSetTypeFromJSONTyped,
    ImageSetTypeToJSON,
} from './ImageSetType';
import type { ProfileStatusType } from './ProfileStatusType';
import {
    ProfileStatusTypeFromJSON,
    ProfileStatusTypeFromJSONTyped,
    ProfileStatusTypeToJSON,
} from './ProfileStatusType';
import type { ProfileTypeAddresses } from './ProfileTypeAddresses';
import {
    ProfileTypeAddressesFromJSON,
    ProfileTypeAddressesFromJSONTyped,
    ProfileTypeAddressesToJSON,
} from './ProfileTypeAddresses';
import type { ProfileTypeComments } from './ProfileTypeComments';
import {
    ProfileTypeCommentsFromJSON,
    ProfileTypeCommentsFromJSONTyped,
    ProfileTypeCommentsToJSON,
} from './ProfileTypeComments';
import type { ProfileTypeEmails } from './ProfileTypeEmails';
import {
    ProfileTypeEmailsFromJSON,
    ProfileTypeEmailsFromJSONTyped,
    ProfileTypeEmailsToJSON,
} from './ProfileTypeEmails';
import type { ProfileTypeRelationships } from './ProfileTypeRelationships';
import {
    ProfileTypeRelationshipsFromJSON,
    ProfileTypeRelationshipsFromJSONTyped,
    ProfileTypeRelationshipsToJSON,
} from './ProfileTypeRelationships';
import type { ProfileTypeRelationshipsSummary } from './ProfileTypeRelationshipsSummary';
import {
    ProfileTypeRelationshipsSummaryFromJSON,
    ProfileTypeRelationshipsSummaryFromJSONTyped,
    ProfileTypeRelationshipsSummaryToJSON,
} from './ProfileTypeRelationshipsSummary';
import type { ProfileTypeTelephones } from './ProfileTypeTelephones';
import {
    ProfileTypeTelephonesFromJSON,
    ProfileTypeTelephonesFromJSONTyped,
    ProfileTypeTelephonesToJSON,
} from './ProfileTypeTelephones';
import type { ProfileTypeType } from './ProfileTypeType';
import {
    ProfileTypeTypeFromJSON,
    ProfileTypeTypeFromJSONTyped,
    ProfileTypeTypeToJSON,
} from './ProfileTypeType';
import type { ProfileTypeURLs } from './ProfileTypeURLs';
import {
    ProfileTypeURLsFromJSON,
    ProfileTypeURLsFromJSONTyped,
    ProfileTypeURLsToJSON,
} from './ProfileTypeURLs';
import type { ReservationStayHistoryFutureInfoType } from './ReservationStayHistoryFutureInfoType';
import {
    ReservationStayHistoryFutureInfoTypeFromJSON,
    ReservationStayHistoryFutureInfoTypeFromJSONTyped,
    ReservationStayHistoryFutureInfoTypeToJSON,
} from './ReservationStayHistoryFutureInfoType';

/**
 * Type provides the detailed information about the profile and its children.
 * @export
 * @interface ProfileType
 */
export interface ProfileType {
    /**
     * 
     * @type {CustomerType}
     * @memberof ProfileType
     */
    customer?: CustomerType;
    /**
     * 
     * @type {CompanyType}
     * @memberof ProfileType
     */
    company?: CompanyType;
    /**
     * 
     * @type {ImageSetType}
     * @memberof ProfileType
     */
    profileImage?: ImageSetType;
    /**
     * 
     * @type {ProfileTypeAddresses}
     * @memberof ProfileType
     */
    addresses?: ProfileTypeAddresses;
    /**
     * Owner Code associated to the profile.
     * @type {string}
     * @memberof ProfileType
     */
    ownerCode?: string;
    /**
     * 
     * @type {ProfileTypeTelephones}
     * @memberof ProfileType
     */
    telephones?: ProfileTypeTelephones;
    /**
     * 
     * @type {ProfileTypeEmails}
     * @memberof ProfileType
     */
    emails?: ProfileTypeEmails;
    /**
     * 
     * @type {ProfileTypeURLs}
     * @memberof ProfileType
     */
    uRLs?: ProfileTypeURLs;
    /**
     * 
     * @type {ProfileTypeComments}
     * @memberof ProfileType
     */
    comments?: ProfileTypeComments;
    /**
     * 
     * @type {ProfileTypeRelationships}
     * @memberof ProfileType
     */
    relationships?: ProfileTypeRelationships;
    /**
     * 
     * @type {ProfileTypeRelationshipsSummary}
     * @memberof ProfileType
     */
    relationshipsSummary?: ProfileTypeRelationshipsSummary;
    /**
     * 
     * @type {ReservationStayHistoryFutureInfoType}
     * @memberof ProfileType
     */
    stayReservationInfoList?: ReservationStayHistoryFutureInfoType;
    /**
     * 
     * @type {ProfileTypeType}
     * @memberof ProfileType
     */
    profileType?: ProfileTypeType;
    /**
     * 
     * @type {ProfileStatusType}
     * @memberof ProfileType
     */
    statusCode?: ProfileStatusType;
    /**
     * Hotel which this profile is registered with. This attribute is not used for configuration.
     * @type {string}
     * @memberof ProfileType
     */
    registeredProperty?: string;
    /**
     * Hotel which this profile is to be registered. This attribute is only used during creation of profile.
     * @type {string}
     * @memberof ProfileType
     */
    requestForHotel?: string;
}

/**
 * Check if a given object implements the ProfileType interface.
 */
export function instanceOfProfileType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileTypeFromJSON(json: any): ProfileType {
    return ProfileTypeFromJSONTyped(json, false);
}

export function ProfileTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'customer': !exists(json, 'customer') ? undefined : CustomerTypeFromJSON(json['customer']),
        'company': !exists(json, 'company') ? undefined : CompanyTypeFromJSON(json['company']),
        'profileImage': !exists(json, 'profileImage') ? undefined : ImageSetTypeFromJSON(json['profileImage']),
        'addresses': !exists(json, 'addresses') ? undefined : ProfileTypeAddressesFromJSON(json['addresses']),
        'ownerCode': !exists(json, 'ownerCode') ? undefined : json['ownerCode'],
        'telephones': !exists(json, 'telephones') ? undefined : ProfileTypeTelephonesFromJSON(json['telephones']),
        'emails': !exists(json, 'emails') ? undefined : ProfileTypeEmailsFromJSON(json['emails']),
        'uRLs': !exists(json, 'uRLs') ? undefined : ProfileTypeURLsFromJSON(json['uRLs']),
        'comments': !exists(json, 'comments') ? undefined : ProfileTypeCommentsFromJSON(json['comments']),
        'relationships': !exists(json, 'relationships') ? undefined : ProfileTypeRelationshipsFromJSON(json['relationships']),
        'relationshipsSummary': !exists(json, 'relationshipsSummary') ? undefined : ProfileTypeRelationshipsSummaryFromJSON(json['relationshipsSummary']),
        'stayReservationInfoList': !exists(json, 'stayReservationInfoList') ? undefined : ReservationStayHistoryFutureInfoTypeFromJSON(json['stayReservationInfoList']),
        'profileType': !exists(json, 'profileType') ? undefined : ProfileTypeTypeFromJSON(json['profileType']),
        'statusCode': !exists(json, 'statusCode') ? undefined : ProfileStatusTypeFromJSON(json['statusCode']),
        'registeredProperty': !exists(json, 'registeredProperty') ? undefined : json['registeredProperty'],
        'requestForHotel': !exists(json, 'requestForHotel') ? undefined : json['requestForHotel'],
    };
}

export function ProfileTypeToJSON(value?: ProfileType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'customer': CustomerTypeToJSON(value.customer),
        'company': CompanyTypeToJSON(value.company),
        'profileImage': ImageSetTypeToJSON(value.profileImage),
        'addresses': ProfileTypeAddressesToJSON(value.addresses),
        'ownerCode': value.ownerCode,
        'telephones': ProfileTypeTelephonesToJSON(value.telephones),
        'emails': ProfileTypeEmailsToJSON(value.emails),
        'uRLs': ProfileTypeURLsToJSON(value.uRLs),
        'comments': ProfileTypeCommentsToJSON(value.comments),
        'relationships': ProfileTypeRelationshipsToJSON(value.relationships),
        'relationshipsSummary': ProfileTypeRelationshipsSummaryToJSON(value.relationshipsSummary),
        'stayReservationInfoList': ReservationStayHistoryFutureInfoTypeToJSON(value.stayReservationInfoList),
        'profileType': ProfileTypeTypeToJSON(value.profileType),
        'statusCode': ProfileStatusTypeToJSON(value.statusCode),
        'registeredProperty': value.registeredProperty,
        'requestForHotel': value.requestForHotel,
    };
}

