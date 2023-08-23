/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
import type { ECertificateType } from './ECertificateType';
import {
    ECertificateTypeFromJSON,
    ECertificateTypeFromJSONTyped,
    ECertificateTypeToJSON,
} from './ECertificateType';
import type { ImageSetType } from './ImageSetType';
import {
    ImageSetTypeFromJSON,
    ImageSetTypeFromJSONTyped,
    ImageSetTypeToJSON,
} from './ImageSetType';
import type { LastStayInfoType } from './LastStayInfoType';
import {
    LastStayInfoTypeFromJSON,
    LastStayInfoTypeFromJSONTyped,
    LastStayInfoTypeToJSON,
} from './LastStayInfoType';
import type { ProfileCashieringType } from './ProfileCashieringType';
import {
    ProfileCashieringTypeFromJSON,
    ProfileCashieringTypeFromJSONTyped,
    ProfileCashieringTypeToJSON,
} from './ProfileCashieringType';
import type { ProfileCommissionType } from './ProfileCommissionType';
import {
    ProfileCommissionTypeFromJSON,
    ProfileCommissionTypeFromJSONTyped,
    ProfileCommissionTypeToJSON,
} from './ProfileCommissionType';
import type { ProfileRestrictions } from './ProfileRestrictions';
import {
    ProfileRestrictionsFromJSON,
    ProfileRestrictionsFromJSONTyped,
    ProfileRestrictionsToJSON,
} from './ProfileRestrictions';
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
import type { ProfileTypePreferenceCollection } from './ProfileTypePreferenceCollection';
import {
    ProfileTypePreferenceCollectionFromJSON,
    ProfileTypePreferenceCollectionFromJSONTyped,
    ProfileTypePreferenceCollectionToJSON,
} from './ProfileTypePreferenceCollection';
import type { ProfileTypeProfileDeliveryMethods } from './ProfileTypeProfileDeliveryMethods';
import {
    ProfileTypeProfileDeliveryMethodsFromJSON,
    ProfileTypeProfileDeliveryMethodsFromJSONTyped,
    ProfileTypeProfileDeliveryMethodsToJSON,
} from './ProfileTypeProfileDeliveryMethods';
import type { ProfileTypeProfileMemberships } from './ProfileTypeProfileMemberships';
import {
    ProfileTypeProfileMembershipsFromJSON,
    ProfileTypeProfileMembershipsFromJSONTyped,
    ProfileTypeProfileMembershipsToJSON,
} from './ProfileTypeProfileMemberships';
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
import type { ProfileTypeRoomOwnershipsList } from './ProfileTypeRoomOwnershipsList';
import {
    ProfileTypeRoomOwnershipsListFromJSON,
    ProfileTypeRoomOwnershipsListFromJSONTyped,
    ProfileTypeRoomOwnershipsListToJSON,
} from './ProfileTypeRoomOwnershipsList';
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
import type { ReservationHistoryFutureInfoType } from './ReservationHistoryFutureInfoType';
import {
    ReservationHistoryFutureInfoTypeFromJSON,
    ReservationHistoryFutureInfoTypeFromJSONTyped,
    ReservationHistoryFutureInfoTypeToJSON,
} from './ReservationHistoryFutureInfoType';
import type { ReservationStayHistoryFutureInfoType } from './ReservationStayHistoryFutureInfoType';
import {
    ReservationStayHistoryFutureInfoTypeFromJSON,
    ReservationStayHistoryFutureInfoTypeFromJSONTyped,
    ReservationStayHistoryFutureInfoTypeToJSON,
} from './ReservationStayHistoryFutureInfoType';
import type { UserDefinedFieldsType } from './UserDefinedFieldsType';
import {
    UserDefinedFieldsTypeFromJSON,
    UserDefinedFieldsTypeFromJSONTyped,
    UserDefinedFieldsTypeToJSON,
} from './UserDefinedFieldsType';

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
     * Trace Code associated to the profile.
     * @type {string}
     * @memberof ProfileType
     */
    traceCode?: string;
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
     * @type {ProfileTypeProfileDeliveryMethods}
     * @memberof ProfileType
     */
    profileDeliveryMethods?: ProfileTypeProfileDeliveryMethods;
    /**
     * 
     * @type {ProfileTypeProfileMemberships}
     * @memberof ProfileType
     */
    profileMemberships?: ProfileTypeProfileMemberships;
    /**
     * 
     * @type {ProfileTypePreferenceCollection}
     * @memberof ProfileType
     */
    preferenceCollection?: ProfileTypePreferenceCollection;
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
     * @type {ReservationHistoryFutureInfoType}
     * @memberof ProfileType
     */
    reservationInfoList?: ReservationHistoryFutureInfoType;
    /**
     * 
     * @type {ReservationStayHistoryFutureInfoType}
     * @memberof ProfileType
     */
    stayReservationInfoList?: ReservationStayHistoryFutureInfoType;
    /**
     * 
     * @type {LastStayInfoType}
     * @memberof ProfileType
     */
    lastStayInfo?: LastStayInfoType;
    /**
     * 
     * @type {ProfileRestrictions}
     * @memberof ProfileType
     */
    profileRestrictions?: ProfileRestrictions;
    /**
     * 
     * @type {ProfileCashieringType}
     * @memberof ProfileType
     */
    cashiering?: ProfileCashieringType;
    /**
     * Contains commission related details for the profile.
     * @type {Array<ProfileCommissionType>}
     * @memberof ProfileType
     */
    commissionInfoList?: Array<ProfileCommissionType>;
    /**
     * 
     * @type {UserDefinedFieldsType}
     * @memberof ProfileType
     */
    userDefinedFields?: UserDefinedFieldsType;
    /**
     * List of e-certificates for the profile.
     * @type {Array<ECertificateType>}
     * @memberof ProfileType
     */
    eCertificates?: Array<ECertificateType>;
    /**
     * Eligible for Fiscal Folio/Payload generation.
     * @type {string}
     * @memberof ProfileType
     */
    eligibleForFiscalFolio?: string;
    /**
     * 
     * @type {ProfileTypeRoomOwnershipsList}
     * @memberof ProfileType
     */
    roomOwnershipsList?: ProfileTypeRoomOwnershipsList;
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
    /**
     * What level this profile is protected.
     * @type {string}
     * @memberof ProfileType
     */
    protectedBy?: string;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof ProfileType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof ProfileType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof ProfileType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof ProfileType
     */
    lastModifierId?: string;
    /**
     * Mark this profile as recently accessed.
     * @type {boolean}
     * @memberof ProfileType
     */
    markAsRecentlyAccessed?: boolean;
    /**
     * "true" setting marks the profile to be kept from being purged, once the profile is marked inactive.
     * @type {boolean}
     * @memberof ProfileType
     */
    markForHistory?: boolean;
    /**
     * Populates true if the profile has commission configured false otherwise.
     * @type {boolean}
     * @memberof ProfileType
     */
    hasCommission?: boolean;
    /**
     * Flag to show inactive Room Owners.
     * @type {boolean}
     * @memberof ProfileType
     */
    showInactiveRoomOwners?: boolean;
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
        'traceCode': !exists(json, 'traceCode') ? undefined : json['traceCode'],
        'ownerCode': !exists(json, 'ownerCode') ? undefined : json['ownerCode'],
        'telephones': !exists(json, 'telephones') ? undefined : ProfileTypeTelephonesFromJSON(json['telephones']),
        'emails': !exists(json, 'emails') ? undefined : ProfileTypeEmailsFromJSON(json['emails']),
        'uRLs': !exists(json, 'uRLs') ? undefined : ProfileTypeURLsFromJSON(json['uRLs']),
        'comments': !exists(json, 'comments') ? undefined : ProfileTypeCommentsFromJSON(json['comments']),
        'profileDeliveryMethods': !exists(json, 'profileDeliveryMethods') ? undefined : ProfileTypeProfileDeliveryMethodsFromJSON(json['profileDeliveryMethods']),
        'profileMemberships': !exists(json, 'profileMemberships') ? undefined : ProfileTypeProfileMembershipsFromJSON(json['profileMemberships']),
        'preferenceCollection': !exists(json, 'preferenceCollection') ? undefined : ProfileTypePreferenceCollectionFromJSON(json['preferenceCollection']),
        'relationships': !exists(json, 'relationships') ? undefined : ProfileTypeRelationshipsFromJSON(json['relationships']),
        'relationshipsSummary': !exists(json, 'relationshipsSummary') ? undefined : ProfileTypeRelationshipsSummaryFromJSON(json['relationshipsSummary']),
        'reservationInfoList': !exists(json, 'reservationInfoList') ? undefined : ReservationHistoryFutureInfoTypeFromJSON(json['reservationInfoList']),
        'stayReservationInfoList': !exists(json, 'stayReservationInfoList') ? undefined : ReservationStayHistoryFutureInfoTypeFromJSON(json['stayReservationInfoList']),
        'lastStayInfo': !exists(json, 'lastStayInfo') ? undefined : LastStayInfoTypeFromJSON(json['lastStayInfo']),
        'profileRestrictions': !exists(json, 'profileRestrictions') ? undefined : ProfileRestrictionsFromJSON(json['profileRestrictions']),
        'cashiering': !exists(json, 'cashiering') ? undefined : ProfileCashieringTypeFromJSON(json['cashiering']),
        'commissionInfoList': !exists(json, 'commissionInfoList') ? undefined : ((json['commissionInfoList'] as Array<any>).map(ProfileCommissionTypeFromJSON)),
        'userDefinedFields': !exists(json, 'userDefinedFields') ? undefined : UserDefinedFieldsTypeFromJSON(json['userDefinedFields']),
        'eCertificates': !exists(json, 'eCertificates') ? undefined : ((json['eCertificates'] as Array<any>).map(ECertificateTypeFromJSON)),
        'eligibleForFiscalFolio': !exists(json, 'eligibleForFiscalFolio') ? undefined : json['eligibleForFiscalFolio'],
        'roomOwnershipsList': !exists(json, 'roomOwnershipsList') ? undefined : ProfileTypeRoomOwnershipsListFromJSON(json['roomOwnershipsList']),
        'profileType': !exists(json, 'profileType') ? undefined : ProfileTypeTypeFromJSON(json['profileType']),
        'statusCode': !exists(json, 'statusCode') ? undefined : ProfileStatusTypeFromJSON(json['statusCode']),
        'registeredProperty': !exists(json, 'registeredProperty') ? undefined : json['registeredProperty'],
        'requestForHotel': !exists(json, 'requestForHotel') ? undefined : json['requestForHotel'],
        'protectedBy': !exists(json, 'protectedBy') ? undefined : json['protectedBy'],
        'createDateTime': !exists(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !exists(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !exists(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !exists(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'markAsRecentlyAccessed': !exists(json, 'markAsRecentlyAccessed') ? undefined : json['markAsRecentlyAccessed'],
        'markForHistory': !exists(json, 'markForHistory') ? undefined : json['markForHistory'],
        'hasCommission': !exists(json, 'hasCommission') ? undefined : json['hasCommission'],
        'showInactiveRoomOwners': !exists(json, 'showInactiveRoomOwners') ? undefined : json['showInactiveRoomOwners'],
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
        'traceCode': value.traceCode,
        'ownerCode': value.ownerCode,
        'telephones': ProfileTypeTelephonesToJSON(value.telephones),
        'emails': ProfileTypeEmailsToJSON(value.emails),
        'uRLs': ProfileTypeURLsToJSON(value.uRLs),
        'comments': ProfileTypeCommentsToJSON(value.comments),
        'profileDeliveryMethods': ProfileTypeProfileDeliveryMethodsToJSON(value.profileDeliveryMethods),
        'profileMemberships': ProfileTypeProfileMembershipsToJSON(value.profileMemberships),
        'preferenceCollection': ProfileTypePreferenceCollectionToJSON(value.preferenceCollection),
        'relationships': ProfileTypeRelationshipsToJSON(value.relationships),
        'relationshipsSummary': ProfileTypeRelationshipsSummaryToJSON(value.relationshipsSummary),
        'reservationInfoList': ReservationHistoryFutureInfoTypeToJSON(value.reservationInfoList),
        'stayReservationInfoList': ReservationStayHistoryFutureInfoTypeToJSON(value.stayReservationInfoList),
        'lastStayInfo': LastStayInfoTypeToJSON(value.lastStayInfo),
        'profileRestrictions': ProfileRestrictionsToJSON(value.profileRestrictions),
        'cashiering': ProfileCashieringTypeToJSON(value.cashiering),
        'commissionInfoList': value.commissionInfoList === undefined ? undefined : ((value.commissionInfoList as Array<any>).map(ProfileCommissionTypeToJSON)),
        'userDefinedFields': UserDefinedFieldsTypeToJSON(value.userDefinedFields),
        'eCertificates': value.eCertificates === undefined ? undefined : ((value.eCertificates as Array<any>).map(ECertificateTypeToJSON)),
        'eligibleForFiscalFolio': value.eligibleForFiscalFolio,
        'roomOwnershipsList': ProfileTypeRoomOwnershipsListToJSON(value.roomOwnershipsList),
        'profileType': ProfileTypeTypeToJSON(value.profileType),
        'statusCode': ProfileStatusTypeToJSON(value.statusCode),
        'registeredProperty': value.registeredProperty,
        'requestForHotel': value.requestForHotel,
        'protectedBy': value.protectedBy,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'markAsRecentlyAccessed': value.markAsRecentlyAccessed,
        'markForHistory': value.markForHistory,
        'hasCommission': value.hasCommission,
        'showInactiveRoomOwners': value.showInactiveRoomOwners,
    };
}

