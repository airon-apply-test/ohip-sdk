/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AddressSearchType } from './AddressSearchType';
import {
    AddressSearchTypeFromJSON,
    AddressSearchTypeFromJSONTyped,
    AddressSearchTypeToJSON,
} from './AddressSearchType';
import type { AnonymizationType } from './AnonymizationType';
import {
    AnonymizationTypeFromJSON,
    AnonymizationTypeFromJSONTyped,
    AnonymizationTypeToJSON,
} from './AnonymizationType';
import type { CountryNameType } from './CountryNameType';
import {
    CountryNameTypeFromJSON,
    CountryNameTypeFromJSONTyped,
    CountryNameTypeToJSON,
} from './CountryNameType';
import type { GuestLastStayInfoType } from './GuestLastStayInfoType';
import {
    GuestLastStayInfoTypeFromJSON,
    GuestLastStayInfoTypeFromJSONTyped,
    GuestLastStayInfoTypeToJSON,
} from './GuestLastStayInfoType';
import type { MembershipInfoType } from './MembershipInfoType';
import {
    MembershipInfoTypeFromJSON,
    MembershipInfoTypeFromJSONTyped,
    MembershipInfoTypeToJSON,
} from './MembershipInfoType';
import type { NameTypeType } from './NameTypeType';
import {
    NameTypeTypeFromJSON,
    NameTypeTypeFromJSONTyped,
    NameTypeTypeToJSON,
} from './NameTypeType';
import type { ResAccompanyGuestInfoType } from './ResAccompanyGuestInfoType';
import {
    ResAccompanyGuestInfoTypeFromJSON,
    ResAccompanyGuestInfoTypeFromJSONTyped,
    ResAccompanyGuestInfoTypeToJSON,
} from './ResAccompanyGuestInfoType';
import type { ResGuestExternalInfoType } from './ResGuestExternalInfoType';
import {
    ResGuestExternalInfoTypeFromJSON,
    ResGuestExternalInfoTypeFromJSONTyped,
    ResGuestExternalInfoTypeToJSON,
} from './ResGuestExternalInfoType';
import type { VIPType } from './VIPType';
import {
    VIPTypeFromJSON,
    VIPTypeFromJSONTyped,
    VIPTypeToJSON,
} from './VIPType';

/**
 * Specifies Company or Travel Agent profile using IATA or Corp. No.
 * @export
 * @interface ResGuestInfoType
 */
export interface ResGuestInfoType {
    /**
     * 
     * @type {MembershipInfoType}
     * @memberof ResGuestInfoType
     */
    membership?: MembershipInfoType;
    /**
     * Salutation of honorific. (e.g., Mr. Mrs., Ms., Miss, Dr.)
     * @type {string}
     * @memberof ResGuestInfoType
     */
    namePrefix?: string;
    /**
     * Given name, first name or names
     * @type {string}
     * @memberof ResGuestInfoType
     */
    givenName?: string;
    /**
     * Alternate given name, first name or names
     * @type {string}
     * @memberof ResGuestInfoType
     */
    alternateGivenName?: string;
    /**
     * The middle name of the person name
     * @type {string}
     * @memberof ResGuestInfoType
     */
    middleName?: string;
    /**
     * e.g "van der", "von", "de"
     * @type {string}
     * @memberof ResGuestInfoType
     */
    surnamePrefix?: string;
    /**
     * Family name, last name.
     * @type {string}
     * @memberof ResGuestInfoType
     */
    surname?: string;
    /**
     * Alternate family name, last name.
     * @type {string}
     * @memberof ResGuestInfoType
     */
    alternateSurname?: string;
    /**
     * Hold various name suffixes and letters (e.g. Jr., Sr., III, Ret., Esq.).
     * @type {string}
     * @memberof ResGuestInfoType
     */
    nameSuffix?: string;
    /**
     * Degree or honors (e.g., Ph.D., M.D.)
     * @type {string}
     * @memberof ResGuestInfoType
     */
    nameTitle?: string;
    /**
     * Full display name
     * @type {string}
     * @memberof ResGuestInfoType
     */
    fullName?: string;
    /**
     * Altername full display name
     * @type {string}
     * @memberof ResGuestInfoType
     */
    alternateFullName?: string;
    /**
     * Phone number
     * @type {string}
     * @memberof ResGuestInfoType
     */
    phoneNumber?: string;
    /**
     * Email address
     * @type {string}
     * @memberof ResGuestInfoType
     */
    email?: string;
    /**
     * Date of birth
     * @type {Date}
     * @memberof ResGuestInfoType
     */
    birthDate?: Date;
    /**
     * Language identification.
     * @type {string}
     * @memberof ResGuestInfoType
     */
    language?: string;
    /**
     * 
     * @type {CountryNameType}
     * @memberof ResGuestInfoType
     */
    nationality?: CountryNameType;
    /**
     * 
     * @type {VIPType}
     * @memberof ResGuestInfoType
     */
    vip?: VIPType;
    /**
     * 
     * @type {AddressSearchType}
     * @memberof ResGuestInfoType
     */
    address?: AddressSearchType;
    /**
     * 
     * @type {AnonymizationType}
     * @memberof ResGuestInfoType
     */
    anonymization?: AnonymizationType;
    /**
     * Collection of accompany guests
     * @type {Array<ResAccompanyGuestInfoType>}
     * @memberof ResGuestInfoType
     */
    accompanyGuests?: Array<ResAccompanyGuestInfoType>;
    /**
     * 
     * @type {ResGuestExternalInfoType}
     * @memberof ResGuestInfoType
     */
    externalInfo?: ResGuestExternalInfoType;
    /**
     * 
     * @type {GuestLastStayInfoType}
     * @memberof ResGuestInfoType
     */
    guestLastStayInfo?: GuestLastStayInfoType;
    /**
     * Guest profile restricted reason code
     * @type {string}
     * @memberof ResGuestInfoType
     */
    guestRestrictedCode?: string;
    /**
     * Guest profile restricted reason description
     * @type {string}
     * @memberof ResGuestInfoType
     */
    guestRestrictedReasonDesc?: string;
    /**
     * True indicates there are restrictions associated with the current profile.
     * @type {boolean}
     * @memberof ResGuestInfoType
     */
    guestRestricted?: boolean;
    /**
     * Unique identifier of the police registration card number.
     * @type {string}
     * @memberof ResGuestInfoType
     */
    registrationCardNo?: string;
    /**
     * 
     * @type {NameTypeType}
     * @memberof ResGuestInfoType
     */
    nameType?: NameTypeType;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof ResGuestInfoType
     */
    id?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof ResGuestInfoType
     */
    type?: string;
}

/**
 * Check if a given object implements the ResGuestInfoType interface.
 */
export function instanceOfResGuestInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResGuestInfoTypeFromJSON(json: any): ResGuestInfoType {
    return ResGuestInfoTypeFromJSONTyped(json, false);
}

export function ResGuestInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResGuestInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'membership': !exists(json, 'membership') ? undefined : MembershipInfoTypeFromJSON(json['membership']),
        'namePrefix': !exists(json, 'namePrefix') ? undefined : json['namePrefix'],
        'givenName': !exists(json, 'givenName') ? undefined : json['givenName'],
        'alternateGivenName': !exists(json, 'alternateGivenName') ? undefined : json['alternateGivenName'],
        'middleName': !exists(json, 'middleName') ? undefined : json['middleName'],
        'surnamePrefix': !exists(json, 'surnamePrefix') ? undefined : json['surnamePrefix'],
        'surname': !exists(json, 'surname') ? undefined : json['surname'],
        'alternateSurname': !exists(json, 'alternateSurname') ? undefined : json['alternateSurname'],
        'nameSuffix': !exists(json, 'nameSuffix') ? undefined : json['nameSuffix'],
        'nameTitle': !exists(json, 'nameTitle') ? undefined : json['nameTitle'],
        'fullName': !exists(json, 'fullName') ? undefined : json['fullName'],
        'alternateFullName': !exists(json, 'alternateFullName') ? undefined : json['alternateFullName'],
        'phoneNumber': !exists(json, 'phoneNumber') ? undefined : json['phoneNumber'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'birthDate': !exists(json, 'birthDate') ? undefined : (new Date(json['birthDate'])),
        'language': !exists(json, 'language') ? undefined : json['language'],
        'nationality': !exists(json, 'nationality') ? undefined : CountryNameTypeFromJSON(json['nationality']),
        'vip': !exists(json, 'vip') ? undefined : VIPTypeFromJSON(json['vip']),
        'address': !exists(json, 'address') ? undefined : AddressSearchTypeFromJSON(json['address']),
        'anonymization': !exists(json, 'anonymization') ? undefined : AnonymizationTypeFromJSON(json['anonymization']),
        'accompanyGuests': !exists(json, 'accompanyGuests') ? undefined : ((json['accompanyGuests'] as Array<any>).map(ResAccompanyGuestInfoTypeFromJSON)),
        'externalInfo': !exists(json, 'externalInfo') ? undefined : ResGuestExternalInfoTypeFromJSON(json['externalInfo']),
        'guestLastStayInfo': !exists(json, 'guestLastStayInfo') ? undefined : GuestLastStayInfoTypeFromJSON(json['guestLastStayInfo']),
        'guestRestrictedCode': !exists(json, 'guestRestrictedCode') ? undefined : json['guestRestrictedCode'],
        'guestRestrictedReasonDesc': !exists(json, 'guestRestrictedReasonDesc') ? undefined : json['guestRestrictedReasonDesc'],
        'guestRestricted': !exists(json, 'guestRestricted') ? undefined : json['guestRestricted'],
        'registrationCardNo': !exists(json, 'registrationCardNo') ? undefined : json['registrationCardNo'],
        'nameType': !exists(json, 'nameType') ? undefined : NameTypeTypeFromJSON(json['nameType']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function ResGuestInfoTypeToJSON(value?: ResGuestInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'membership': MembershipInfoTypeToJSON(value.membership),
        'namePrefix': value.namePrefix,
        'givenName': value.givenName,
        'alternateGivenName': value.alternateGivenName,
        'middleName': value.middleName,
        'surnamePrefix': value.surnamePrefix,
        'surname': value.surname,
        'alternateSurname': value.alternateSurname,
        'nameSuffix': value.nameSuffix,
        'nameTitle': value.nameTitle,
        'fullName': value.fullName,
        'alternateFullName': value.alternateFullName,
        'phoneNumber': value.phoneNumber,
        'email': value.email,
        'birthDate': value.birthDate === undefined ? undefined : (value.birthDate.toISOString().substr(0,10)),
        'language': value.language,
        'nationality': CountryNameTypeToJSON(value.nationality),
        'vip': VIPTypeToJSON(value.vip),
        'address': AddressSearchTypeToJSON(value.address),
        'anonymization': AnonymizationTypeToJSON(value.anonymization),
        'accompanyGuests': value.accompanyGuests === undefined ? undefined : ((value.accompanyGuests as Array<any>).map(ResAccompanyGuestInfoTypeToJSON)),
        'externalInfo': ResGuestExternalInfoTypeToJSON(value.externalInfo),
        'guestLastStayInfo': GuestLastStayInfoTypeToJSON(value.guestLastStayInfo),
        'guestRestrictedCode': value.guestRestrictedCode,
        'guestRestrictedReasonDesc': value.guestRestrictedReasonDesc,
        'guestRestricted': value.guestRestricted,
        'registrationCardNo': value.registrationCardNo,
        'nameType': NameTypeTypeToJSON(value.nameType),
        'id': value.id,
        'type': value.type,
    };
}

